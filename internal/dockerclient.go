package internal

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

type DockerAgent struct {
	Client *client.Client
}

func DockerClient() (*DockerAgent, error) {
	apiClient, err := client.NewClientWithOpts(
		client.FromEnv,
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		panic(err)
	}
	defer apiClient.Close()

	return &DockerAgent{Client: apiClient}, nil

}

func (da DockerAgent) ListOfContainers() error {
	containers, err := da.Client.ContainerList(context.Background(), container.ListOptions{All: true})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("ID: %s, Image: %s, Status: %s\n", ctr.ID, ctr.Image, ctr.Status)
	}
	return nil
}
