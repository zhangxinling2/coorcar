package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestCreateTrip(t *testing.T) {
	testCases := []struct {
		name       string
		accountId  string
		tripId     string
		tripStatus rentalpb.TripStatus
		wantErr    bool
	}{
		{
			name:       "finished",
			accountId:  "account1",
			tripId:     "65a713221b886d3a9bba21e3",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "another_finished",
			accountId:  "account1",
			tripId:     "65a713221b886d3a9bba21e4",
			tripStatus: rentalpb.TripStatus_FINISHED,
		},
		{
			name:       "progressing",
			accountId:  "account1",
			tripId:     "65a713221b886d3a9bba21e5",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
		{
			name:       "progressing",
			accountId:  "account1",
			tripId:     "65a713221b886d3a9bba21e6",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantErr:    true,
		},
		{
			name:       "another account progressing",
			accountId:  "account2",
			tripId:     "65a713221b886d3a9bba21e7",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
		},
	}
	c := context.Background()
	mo, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	m := NewMongo(mo.Database("coolcar"))
	err = mongotesting.CreateIndexes(c, mo.Database("coolcar"))
	if err != nil {
		t.Fatalf("cannot create index %v", err)
	}
	for _, tc := range testCases {
		mgo.NewObjId = func() primitive.ObjectID {
			return objid.MustFromId(id.TripId(tc.tripId))
		}
		tr, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: tc.accountId,
			Status:    tc.tripStatus,
		})
		if tc.wantErr {
			if err == nil {
				t.Errorf("error excepted;got none")

			}
			continue
		}
		if err != nil {
			t.Errorf("cannot create trip %v", err)
			continue
		}
		if tr.ID.Hex() != tc.tripId {
			t.Errorf("%s incorrec  tripid;want:%q;got:%q", tc.name, tc.tripId, tr.ID.Hex())
		}
	}
}
func TestGetTrip(t *testing.T) {
	c := context.Background()
	mo, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	m := NewMongo(mo.Database("coolcar"))
	res, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: "account1",
		CarId:     "car123",
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  10.11,
				Longitude: 12.11,
			},
			KmDriven: 0,
			PoiName:  "startpoint",
		},
		End: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  10.11,
				Longitude: 12.11,
			},
			KmDriven: 12,
			PoiName:  "endpoint",
		},
		Status: rentalpb.TripStatus_FINISHED,
	})
	if err != nil {
		t.Fatalf("create trip failed:%v", err)
	}

	tr, err := m.GetTrip(c, objid.ToTripId(res.ID), "account1")
	if err != nil {
		t.Fatalf("get trip failed:%v", err)
	}
	if diff := cmp.Diff(res, tr, protocmp.Transform()); diff != "" {
		t.Errorf("result differs: -want -got:%s", diff)
	}
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.NewWithMongoDocker(m))
}
