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
		Start: &rentalpb.LocationStatus{
			Location: &rentalpb.Location{
				Latitude:  32.123,
				Longitude: 114.2525,
			},
		},
	}
	goldenRes := `{"account_id":"account1","car_id":"car1","start":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔"},"current":{"location":{"latitude":32.123,"longitude":114.2525},"poi_name":"广州塔"},"identity_id":"identity1"}`

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
	ID  id.IndetityId
	Err error
}

func (p *profileManager) Verify(ctx context.Context, id id.AccountId) (id.IndetityId, error) {
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

func TestMain(m *testing.M) {
	os.Exit(mongotesting.NewWithMongoDocker(m))
}
