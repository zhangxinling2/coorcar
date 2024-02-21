package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
)

func main() {
	dc, err := client.NewClientWithOpts()
	if err != nil {
		log.Fatalf("fail new client %v", err)
	}
	c := context.Background()
	res, err := dc.ContainerCreate(c, &container.Config{
		Image: "mongo",
		ExposedPorts: nat.PortSet{
			"27017/tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
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
	err = dc.ContainerStart(c, res.ID, types.ContainerStartOptions{})
	if err != nil {
		log.Fatalf("fail start container %v", err)
	}
	inspRes, err := dc.ContainerInspect(c, res.ID)
	if err != nil {
		log.Fatalf("fail inspect container %v", err)
	}
	fmt.Printf("port map %v", inspRes.NetworkSettings.Ports["27017/tcp"])
	time.Sleep(3 * time.Second)
	dc.ContainerRemove(c, res.ID, types.ContainerRemoveOptions{Force: true})
}
