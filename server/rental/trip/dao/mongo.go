package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField      = "trip"
	accountIdField = tripField + ".accountid"
	statusField    = tripField + ".status"
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

func (m *Mongo) GetTrips(ctx context.Context, accountId id.AccountId, status rentalpb.TripStatus) ([]*TripRecord, error) {
	filter := bson.M{
		accountIdField: accountId,
	}
	if status != rentalpb.TripStatus_TS_NOT_SPECIFIED {
		filter[statusField] = status
	}
	c, err := m.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var trips []*TripRecord
	for c.Next(ctx) {
		var trip TripRecord
		err = c.Decode(&trip)
		if err != nil {
			return nil, err
		}
		trips = append(trips, &trip)
	}
	return trips, nil
}
func (m *Mongo) UpdateTrip(ctx context.Context, id id.TripId, aid id.AccountId, updatedAt int64, trip *rentalpb.Trip) error {
	objId, err := objid.FromId(id)
	if err != nil {
		return err
	}
	newUpdatedAt := mgo.NewUpdatedAt()
	res, err := m.col.UpdateOne(ctx, bson.M{
		mgo.IDFieldName:        objId,
		accountIdField:         aid.String(),
		mgo.UpdatedAtFieldName: updatedAt,
	}, mgo.Set(bson.M{
		tripField:              trip,
		mgo.UpdatedAtFieldName: newUpdatedAt,
	}),
	)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return fmt.Errorf("no match record")
	}
	return nil
}
