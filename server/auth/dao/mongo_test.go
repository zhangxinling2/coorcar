package dao

import (
	"context"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	mongotesting "coolcar/shared/mongo/testing"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestResolveAccountID(t *testing.T) {
	testcases := []struct {
		name   string
		openId string
		want   string
		res    id.AccountId
		err    error
	}{
		{
			name:   "existing_1",
			openId: "123",
			want:   "65a713221b886d3a9bba21e3",
		},
		{
			name:   "existing_2",
			openId: "456",
			want:   "65a713221b886d3a9bba21e4",
		},
		{
			name:   "not_existing",
			openId: "789",
			want:   "65a713221b886d3a9bba21e5",
		},
	}
	c := context.Background()
	mo, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	m := NewMongo(mo.Database("coolcar"))
	_, err = m.col.InsertMany(c, []interface{}{
		bson.M{
			openIdField:     "123",
			mgo.IDFieldName: objid.MustFromId(id.AccountId("65a713221b886d3a9bba21e3")),
		},
		bson.M{
			openIdField:     "456",
			mgo.IDFieldName: objid.MustFromId(id.AccountId("65a713221b886d3a9bba21e4")),
		},
	})
	if err != nil {
		panic(err)
	}
	mgo.NewObjId = func() primitive.ObjectID {
		return objid.MustFromId(id.AccountId("65a713221b886d3a9bba21e5"))
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			tc.res, err = m.ResolveAccountID(c, tc.openId)
			if err != nil {
				t.Errorf("failed resolve %v", err)
			} else {
				if tc.res.String() != tc.want {
					t.Errorf("resolve error id: %q want :%q", tc.res, tc.want)
				}
			}
		})
	}

}
func NewObjId(id string) primitive.ObjectID {
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		panic(err)
	}
	return objId
}
func TestMain(m *testing.M) {
	os.Exit(mongotesting.NewWithMongoDocker(m))
}
