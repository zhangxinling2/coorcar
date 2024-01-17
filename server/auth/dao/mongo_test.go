package dao

import (
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestResolveAccountID(t *testing.T) {
	c := context.Background()
	mo, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017/coolcar?readPreference=primary&ssl=false"))
	if err != nil {
		t.Fatalf("cannot connect %v", err)
	}
	m := NewMongo(mo.Database("coolcar"))
	id, err := m.ResolveAccountID(c, "123")
	if err != nil {
		t.Errorf("failed resolve %v", err)
	} else {
		want := "65a713221b886d3a9bba21e3"
		if id != want {
			t.Errorf("resolve error id: %q want :%q", id, want)
		}
	}

}
