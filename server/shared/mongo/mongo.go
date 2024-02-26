package mgo

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	IDFieldName        = "_id"
	UpdatedAtFieldName = "updatedAt"
)

var NewObjId = primitive.NewObjectID
var NewUpdatedAt = func() int64 {
	return time.Now().UnixNano()
}

type ObjIdField struct {
	ID primitive.ObjectID `bson:"_id"`
}
type UpdatedAtField struct {
	UpdatedAt int64 `bson:"updatedat"`
}

func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}
