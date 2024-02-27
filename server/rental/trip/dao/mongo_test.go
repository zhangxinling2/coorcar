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

func TestGetTrips(t *testing.T) {
	readyCases := []struct {
		name       string
		accountId  string
		tripId     string
		tripStatus rentalpb.TripStatus
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
	}
	c := context.Background()
	client, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	m := NewMongo(client.Database("coolcar"))
	for _, rc := range readyCases {
		mgo.NewObjIdWithValue(id.TripId(rc.tripId))
		_, err := m.CreateTrip(c, &rentalpb.Trip{
			AccountId: rc.accountId,
			Status:    rc.tripStatus,
		})
		if err != nil {
			t.Fatalf("cannot create trip:%v", err)
		}
	}
	testCases := []struct {
		name       string
		accountId  string
		wantCount  int
		tripStatus rentalpb.TripStatus
	}{
		{
			name:       "account1_all",
			accountId:  "account1",
			tripStatus: rentalpb.TripStatus_TS_NOT_SPECIFIED,
			wantCount:  3,
		},
		{
			name:       "account1_completed",
			accountId:  "account1",
			tripStatus: rentalpb.TripStatus_FINISHED,
			wantCount:  2,
		}, {
			name:       "account1_in_progressing",
			accountId:  "account1",
			tripStatus: rentalpb.TripStatus_IN_PROGRESS,
			wantCount:  1,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := m.GetTrips(c, id.AccountId(tc.accountId), tc.tripStatus)
			if err != nil {
				t.Fatalf("Get trips error:%v", err)
			}
			if len(res) != tc.wantCount {
				t.Fatalf("get num error want %d got %d", tc.wantCount, len(res))
			}
		})
	}
}
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
func TestUpdateTrip(t *testing.T) {
	c := context.Background()
	mo, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	now := int64(10000)
	mgo.NewUpdatedAt = func() int64 {
		return int64(now)
	}
	mgo.NewObjId = func() primitive.ObjectID {
		return objid.MustFromId(id.TripId("65a713221b886d3a9bba21e7"))
	}
	m := NewMongo(mo.Database("coolcar"))
	res, err := m.CreateTrip(c, &rentalpb.Trip{
		AccountId: "account1",
		Start: &rentalpb.LocationStatus{
			PoiName: "startpoint",
		},
	})
	if err != nil {
		t.Fatalf("create trip failed:%v", err)
	}
	res.Trip.Start.PoiName = "startpoint_update"
	updatedTrip := res.Trip
	testcases := []struct {
		name      string
		wantErr   bool
		now       int64
		updatedAt int64
	}{
		{
			name:      "normal_update",
			now:       20000,
			updatedAt: 10000,
		},
		{
			name:      "updated_with_stale_timestamp",
			now:       30000,
			updatedAt: 10000,
			wantErr:   true,
		},
		{
			name:      "update_with_refetch",
			now:       40000,
			updatedAt: 20000,
		},
	}
	for _, tc := range testcases {
		now = tc.now
		err := m.UpdateTrip(c, id.TripId("65a713221b886d3a9bba21e7"), id.AccountId("account1"), tc.updatedAt, updatedTrip)
		if tc.wantErr {
			if err == nil {
				t.Errorf("%s:want err;got nil", tc.name)
			} else {
				continue
			}

		} else {
			if err != nil {
				t.Errorf("%s:want nil;got err:%v", tc.name, err)
			}
		}
		if err != nil {
			t.Errorf("%s:want nil;got err:%v", tc.name, err)
		}
		updatedTrip, err := m.GetTrip(c, id.TripId("65a713221b886d3a9bba21e7"), id.AccountId("account1"))
		if err != nil {
			t.Fatalf("%s cannot get trip after update %v", tc.name, err)
		}
		if tc.now != updatedTrip.UpdatedAt {
			t.Errorf("%s:incorrect updatedat want:%d,got:%d", tc.name, tc.now, updatedTrip.UpdatedAt)
		}
	}
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.NewWithMongoDocker(m))
}
