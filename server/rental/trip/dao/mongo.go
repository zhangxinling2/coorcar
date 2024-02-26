package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField      = "trip"
	accountIdField = tripField + ".accountid"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

type TripRecord struct {
	mgo.ObjIdField     `bson:"inline"`
	mgo.UpdatedAtField `bson:"inline"`
	Trip               *rentalpb.Trip `bson:"trip"`
}

func (m *Mongo) CreateTrip(ctx context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	t := &TripRecord{
		Trip: trip,
	}
	t.ID = mgo.NewObjId()
	t.UpdatedAt = mgo.NewUpdatedAt()
	_, err := m.col.InsertOne(ctx, t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (m *Mongo) GetTrip(ctx context.Context, id id.TripId, accountId id.AccountId) (*TripRecord, error) {
	objId, err := objid.FromId(id)
	if err != nil {
		return nil, err
	}
	res := m.col.FindOne(ctx, bson.M{
		mgo.IDFieldName: objId,
		accountIdField:  accountId,
	})
	var tr TripRecord
	err = res.Decode(&tr)
	if err != nil {
		return nil, err
	}
	return &tr, nil
}
