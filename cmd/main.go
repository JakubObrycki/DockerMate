package main

import (
	"docker-orch/internal"
	"fmt"
	"os"
)

func main() {
	agent, err := internal.DockerClient()
	if err != nil {
		fmt.Println("Docker client creation error:", err)
		os.Exit(1)
	}

	err = agent.ListOfContainers()
	if err != nil {
		fmt.Println("Error downloading containers:", err)
		os.Exit(1)
	}
}
