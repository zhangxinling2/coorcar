package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	profileField   = "profile"
	accountIdField = "accountid"
	statusField    = profileField + ".identitystatus"
)

type Mongo struct {
	col *mongo.Collection
}
type ProfileRecord struct {
	AccountId string            `bson:"accountid"`
	Profile   *rentalpb.Profile `bson:"profile"`
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("profile"),
	}
}

func (m *Mongo) GetProfile(ctx context.Context, aid id.AccountId) (*rentalpb.Profile, error) {
	sr := m.col.FindOne(ctx, byAccountId(aid))
	if err := sr.Err(); err != nil {
		return nil, err
	}
	var pr ProfileRecord
	err := sr.Decode(&pr)
	if err != nil {
		return nil, err
	}
	return pr.Profile, nil
}
func (m *Mongo) UpdateProfile(ctx context.Context, aid id.AccountId, prevStatus rentalpb.IdentityStatus, p *rentalpb.Profile) error {
	_, err := m.col.UpdateOne(ctx, bson.M{
		accountIdField: aid.String(),
		statusField:    prevStatus,
	}, mgo.Set(bson.M{
		profileField: p,
	}), options.Update().SetUpsert(true))
	if err != nil {
		return err
	}
	return nil
}
func byAccountId(aid id.AccountId) bson.M {
	return bson.M{
		accountIdField: aid.String(),
	}
}
