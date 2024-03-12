package profile

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/rental/profile/dao"
	"coolcar/shared/auth"
	"coolcar/shared/id"
	mongotesting "coolcar/shared/mongo/testing"
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
	s.Logger = z
	aid := "account1"
	c := context.Background()
	c = auth.ContextWithAccountId(c, id.AccountId(aid))
	client, err := mongotesting.NewClient(c)
	if err != nil {
		t.Error("cannot create mongoclient")
	}
	db := client.Database("coolcar")
	s.Mongo = dao.NewMongo(db)
	mongotesting.CreateIndexes(c, db)
	testCases := []struct {
		name       string
		op         func() (*rentalpb.Profile, error)
		wantStatus rentalpb.IdentityStatus
		wantErr    bool
	}{
		{
			name: "get null profile",
			op: func() (*rentalpb.Profile, error) {
				req := &rentalpb.GetProfileRequest{}
				pr, err := s.GetProfile(c, req)
				if err != nil {
					return nil, err
				}
				return pr, nil
			},
			wantStatus: rentalpb.IdentityStatus_UNSUBMITTED,
		},
		{
			name: "submit profile",
			op: func() (*rentalpb.Profile, error) {
				identity := &rentalpb.Identity{
					Name:            "account1",
					Gender:          rentalpb.Gender_MALE,
					LicNumber:       "1234567",
					BirthDataMillis: 1924853,
				}
				pr, err := s.SubmitProfile(c, identity)
				if err != nil {
					return nil, err
				}
				return pr, nil
			},
			wantStatus: rentalpb.IdentityStatus_PENDING,
		},
		{
			name: "submit profile again",
			op: func() (*rentalpb.Profile, error) {
				identity := &rentalpb.Identity{
					Name:            "account1",
					Gender:          rentalpb.Gender_MALE,
					LicNumber:       "1234567",
					BirthDataMillis: 1924853,
				}
				pr, err := s.SubmitProfile(c, identity)
				if err != nil {
					return nil, err
				}
				return pr, nil
			},
			wantErr: true,
		},
		{
			name: "get profile",
			op: func() (*rentalpb.Profile, error) {
				req := &rentalpb.GetProfileRequest{}
				pr, err := s.GetProfile(c, req)
				if err != nil {
					return nil, err
				}
				return pr, nil
			},
			wantStatus: rentalpb.IdentityStatus_PENDING,
		},
		{
			name: "clear profile",
			op: func() (*rentalpb.Profile, error) {
				req := &rentalpb.GetProfileRequest{}
				pr, err := s.GetProfile(c, req)
				if err != nil {
					return nil, err
				}
				pr.IdentityStatus = rentalpb.IdentityStatus_VERIFIED
				err = s.Mongo.UpdateProfile(c, id.AccountId(aid), rentalpb.IdentityStatus_PENDING, pr)
				if err != nil {
					return nil, err
				}
				req2 := &rentalpb.ClearProfileRequest{}
				pr, err = s.ClearProfile(c, req2)
				if err != nil {
					return nil, err
				}
				return pr, nil
			},
			wantStatus: rentalpb.IdentityStatus_UNSUBMITTED,
		},
	}
	for _, tc := range testCases {
		pro, err := tc.op()
		if tc.wantErr {
			if err == nil {
				t.Errorf("%s want err ;got nil", tc.name)
			} else {
				continue
			}
		}
		if err != nil {
			t.Errorf("cannot op %s func", tc.name)
		}

		if tc.wantStatus != pro.IdentityStatus {
			t.Errorf("want status %q;got status %q", tc.wantStatus, pro.IdentityStatus)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.NewWithMongoDocker(m))
}
