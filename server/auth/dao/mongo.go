package dao

import (
	"context"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIdField = "open_id"

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}
func (m *Mongo) ResolveAccountID(c context.Context, openid string) (id.AccountId, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIdField: openid,
	}, mgo.SetOnInsert(bson.M{
		mgo.IDFieldName: mgo.NewObjId(),
		openIdField:     openid,
	}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("canot FindOneAndUpdate %v", err)
	}
	var row mgo.ObjIdField
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("canot Decode %v", err)
	}
	return objid.ToAccountId(row.ID), nil
}
