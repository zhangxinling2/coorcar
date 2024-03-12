package mongotesting

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	image           = "mongo"
	containerPort   = "27017/tcp"
	defaultMongoURI = "mongodb://localhost:27017"
)

var mongoURI string

func NewWithMongoDocker(m *testing.M) int {
	dc, err := client.NewClientWithOpts()
	if err != nil {
		log.Fatalf("fail new client %v", err)
	}
	c := context.Background()
	res, err := dc.ContainerCreate(c, &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", //会自动寻找可用端口
				},
			},
		},
	}, nil, nil, "mongotest")
	if err != nil {
		log.Fatalf("fail new container %v", err)
	}
	defer func() {
		err := dc.ContainerRemove(c, res.ID, types.ContainerRemoveOptions{Force: true})
		if err != nil {
			panic(err)
		}
	}()
	err = dc.ContainerStart(c, res.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Fatalf("fail start container %v", err)
	}
	inspRes, err := dc.ContainerInspect(c, res.ID)
	if err != nil {
		log.Fatalf("fail inspect container %v", err)
	}
	hostPort := inspRes.NetworkSettings.Ports[containerPort][0]
	mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)
	return m.Run()
}
func NewClient(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
}
func NewDefaultClient(ctx context.Context) (*mongo.Client, error) {
	return mongo.Connect(ctx, options.Client().ApplyURI(defaultMongoURI))
}
func CreateIndexes(ctx context.Context, db *mongo.Database) error {
	_, err := db.Collection("auth").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "open_id", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}
	_, err = db.Collection("trip").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "trip.accountid", Value: 1},
			{Key: "trip.status", Value: 1},
		},
		Options: options.Index().SetUnique(true).SetPartialFilterExpression(bson.M{
			"trip.status": 1,
		}),
	})
	if err != nil {
		return err
	}
	_, err = db.Collection("profile").Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{
			{Key: "accountid", Value: 1},
		},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return err
	}
	return nil
}
