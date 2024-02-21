package dao

import (
	"context"
	mgo "coolcar/shared/mongo"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const openIdField = "open_id"

type Mongo struct {
	col   *mongo.Collection
	objId func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:   db.Collection("account"),
		objId: primitive.NewObjectID,
	}
}
func (m *Mongo) ResolveAccountID(c context.Context, openid string) (string, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIdField: openid,
	}, mgo.SetOnInsert(bson.M{
		mgo.IDField: m.objId(),
		openIdField: openid,
	}),
		options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("canot FindOneAndUpdate %v", err)
	}
	var row mgo.ObjId
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("canot Decode %v", err)
	}
	return row.ID.Hex(), nil
}