package dao

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}
func (m *Mongo) ResolveAccountID(c context.Context, openid string) (string, error) {
	res := m.col.FindOneAndUpdate(c, bson.M{
		"open_id": openid,
	}, bson.M{
		"$set": bson.M{
			"open_id": openid,
		},
	}, options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", fmt.Errorf("canot FindOneAndUpdate %v", err)
	}
	var row struct {
		ID primitive.ObjectID `bson:"_id"`
	}
	err := res.Decode(&row)
	if err != nil {
		return "", fmt.Errorf("canot Decode %v", err)
	}
	return row.ID.Hex(), nil
}
