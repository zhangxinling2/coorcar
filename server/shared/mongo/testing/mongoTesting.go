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
)

const (
	image         = "mongo"
	containerPort = "27017/tcp"
)

func NewWithMongoDocker(m *testing.M, mongoURI *string) int {
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
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)
	return m.Run()
}
