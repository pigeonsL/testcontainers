package docker_compose

import (
	"context"
	"fmt"
	"github.com/testcontainers/testcontainers-go"
	"log"
)

func CreateParallelContainers() {
	ctx := context.Background()

	requests := testcontainers.ParallelContainerRequest{
		{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "nginx",
				ExposedPorts: []string{
					"10080/tcp",
				},
			},
			Started: true,
		},
		{
			ContainerRequest: testcontainers.ContainerRequest{
				Image: "nginx",
				ExposedPorts: []string{
					"10081/tcp",
				},
			},
			Started: true,
		},
	}

	res, err := testcontainers.ParallelContainers(ctx, requests, testcontainers.ParallelContainersOptions{})

	if err != nil {
		e, ok := err.(testcontainers.ParallelContainersError)
		if !ok {
			log.Fatalf("unknown error: %v", err)
		}

		for _, pe := range e.Errors {
			fmt.Println(pe.Request, pe.Error)
		}
		return
	}

	for _, c := range res {
		defer c.Terminate(ctx)
	}
}
