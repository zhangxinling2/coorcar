package trip

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/trip/client/poi"
	"coolcar/rental/trip/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	mongotesting "coolcar/shared/mongo/testing"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"go.uber.org/zap"
)

func TestCreateTrip(t *testing.T) {
	s := &Service{}
	z, err := zap.NewDevelopment()
	if err != nil {
		t.Error("cannot create zap log", err.Error())
	}
	pm := &profileManager{
		ID: "identity1",
	}
	cm := &carManager{}
	pom := &poi.Manager{}
	s.ProfileManager = pm
	s.CarManager = cm
	s.PoiManager = pom
	s.Logger = z
	c := context.Background()
	c = auth.ContextWithAccountId(c, "account1")
	client, err := mongotesting.NewClient(c)
	if err != nil {
		t.Error("cannot create mongoclient")
	}
	s.Mongo = dao.NewMongo(client.Database("coolcar"))
	req := &rentalpb.CreateTripRequest{
		CarId: "car1",
		Start: &rentalpb.Location{
			Latitude:  32.123,
			Longitude: 114.2525,
		},
	}
	goldenRes := `{"account_id":"account1","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔","timestamp_sec":1709556870},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔","timestamp_sec":1709556870},"identity_id":"identity1"}`
	nowFunc = func() int64 {
		return 1709556870
	}
	testCases := []struct {
		name           string
		tripId         string
		profileError   error
		carVerifyError error
		carUnlockError error
		want           string
		wantErr        bool
	}{
		{
			name:   "normal_req",
			tripId: "65a713221b886d3a9bba21e5",
			want:   goldenRes,
		},
		{
			name:         "profileManager_error",
			tripId:       "65a713221b886d3a9bba21e6",
			wantErr:      true,
			profileError: fmt.Errorf("profile"),
		},
		{
			name:           "carVerify_error",
			tripId:         "65a713221b886d3a9bba21e7",
			wantErr:        true,
			carVerifyError: fmt.Errorf("carVerify"),
		},
		{
			name:           "carUnlock_error",
			tripId:         "65a713221b886d3a9bba21e8",
			carUnlockError: fmt.Errorf("carUnlock"),
			want:           goldenRes,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mgo.NewObjIdWithValue(id.TripId(tc.tripId))
			pm.Err = tc.profileError
			cm.VerifyErr = tc.carVerifyError
			cm.UnlockErr = tc.carUnlockError
			te, err := s.CreateTrip(c, req)
			if tc.wantErr {
				if err == nil {
					t.Errorf("want err ;got nil")
				} else {
					return
				}
			}
			if err != nil {
				t.Errorf("connot create trip:%v", err)
			}
			b, err := json.Marshal(te.Trip)
			if err != nil {
				t.Errorf("cannot marshal to json:%v", err)
			}
			tripStr := string(b)
			if tc.want != tripStr {
				t.Errorf("incorrect res,want:%s,got:%s", tc.want, tripStr)
			}
		})
	}
}

type profileManager struct {
	ID  id.IdentityId
	Err error
}

func (p *profileManager) Verify(ctx context.Context, id id.AccountId) (id.IdentityId, error) {
	return p.ID, p.Err
}

type carManager struct {
	VerifyErr error
	UnlockErr error
}

func (c *carManager) Verify(ctx context.Context, id id.CarId, loc *rentalpb.Location) error {
	return c.VerifyErr
}
func (c *carManager) UnLock(ctx context.Context, id id.CarId) error {
	return c.UnlockErr
}
func TestLifeCycle(t *testing.T) {
	c := context.Background()
	c = auth.ContextWithAccountId(c, id.AccountId("account_for_lifecycle"))
	s := newService(c, t, &profileManager{}, &carManager{})
	tid := id.TripId("5f8132eb22714bf629489056")
	mgo.NewObjIdWithValue(tid)
	testCases := []struct {
		name    string
		now     int64
		want    string
		wantErr bool
		op      func() (*rentalpb.Trip, error)
	}{

		{
			name: "create_trip",
			now:  10000,
			op: func() (*rentalpb.Trip, error) {
				e, err := s.CreateTrip(c, &rentalpb.CreateTripRequest{
					CarId: "car1",
					Start: &rentalpb.Location{
						Latitude:  32.123,
						Longitude: 114.2525,
					},
				})
				if err != nil {
					return nil, err
				}
				return e.Trip, nil
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔","timestamp_sec":10000},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔","timestamp_sec":10000},"status":1}`,
		},
		{
			name: "update_trip",
			now:  20000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
					Current: &rentalpb.Location{
						Latitude:  28.234234,
						Longitude: 123.243255,
					},
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔","timestamp_sec":10000},"current":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":10677,"km_driven":110.22457407124267,"poi_name":"广州塔","timestamp_sec":20000},"status":1}`,
		},
		{
			name: "finish_trip",
			now:  30000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id:      tid.String(),
					EndTrip: true,
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔","timestamp_sec":10000},"current":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":18853,"km_driven":215.53264623224982,"poi_name":"广州塔","timestamp_sec":30000},"end":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":18853,"km_driven":215.53264623224982,"poi_name":"广州塔","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "query_trip",
			now:  40000,
			op: func() (*rentalpb.Trip, error) {
				return s.GetTrip(c, &rentalpb.GetTripRequest{
					Id: tid.String(),
				})
			},
			want: `{"account_id":"account_for_lifecycle","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔","timestamp_sec":10000},"current":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":18853,"km_driven":215.53264623224982,"poi_name":"广州塔","timestamp_sec":30000},"end":{"location":{"latitude":28.234234,"longitude":123.243255},"fee_cent":18853,"km_driven":215.53264623224982,"poi_name":"广州塔","timestamp_sec":30000},"status":2}`,
		},
		{
			name: "update_after_finished",
			now:  50000,
			op: func() (*rentalpb.Trip, error) {
				return s.UpdateTrip(c, &rentalpb.UpdateTripRequest{
					Id: tid.String(),
				})
			},
			wantErr: true,
		},
	}
	rand.Seed(1345)
	for _, cc := range testCases {
		nowFunc = func() int64 {
			return cc.now
		}
		trip, err := cc.op()
		if cc.wantErr {
			if err == nil {
				t.Errorf("%s: want error; got none", cc.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("%s: operation failed: %v", cc.name, err)
			continue
		}
		b, err := json.Marshal(trip)
		if err != nil {
			t.Errorf("%s: failed marshalling response: %v", cc.name, err)
		}
		got := string(b)
		if cc.want != got {
			t.Errorf("%s: incorrect response; want: %s, got: %s", cc.name, cc.want, got)
		}
	}
}
func newService(c context.Context, t *testing.T, pm ProfileManager, cm CarManager) *Service {
	z, err := zap.NewDevelopment()
	if err != nil {
		t.Error("cannot create zap log", err.Error())
	}
	client, err := mongotesting.NewClient(c)
	if err != nil {
		t.Error("cannot create mongoclient")
	}
	db := client.Database("coolcar")
	mongo := dao.NewMongo(db)
	mongotesting.CreateIndexes(c, db)
	return &Service{
		Logger:         z,
		Mongo:          mongo,
		ProfileManager: pm,
		CarManager:     cm,
		PoiManager:     &poi.Manager{},
	}
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.NewWithMongoDocker(m))
}
