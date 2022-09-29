package docker_compose

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"strings"
)

var compose *testcontainers.LocalDockerCompose

func DockerComposeStart() {
	composeFilePaths := []string{"resource/docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose = testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
		WithCommand([]string{"up", "-d"}).
		Invoke()
	err := execError.Error
	if err != nil {
		println(fmt.Errorf("could not run compose file: %v - %v", composeFilePaths, err))
	}
}

func DockerComposeDown() {
	if compose != nil {
		execError := compose.Down()
		if execError.Error != nil {
			println(fmt.Errorf("failed when running: %v", execError.Command))
		}
	} else {
		println("failed to down compose because compose is nil")
	}
}

func DockerComposeRestart(serviceName string) {
	if compose != nil {
		execError := compose.WithCommand([]string{"restart", serviceName}).Invoke()
		if execError.Error != nil {
			println(fmt.Errorf("failed when running: %v", execError.Command))
		}
	} else {
		println("failed to down compose because compose is nil")
	}
}

func ShowDockerComposeService() {
	if compose != nil {
		execError := compose.WithCommand([]string{"ps", "-q"}).Invoke()
		if execError.Error != nil {
			println(fmt.Errorf("failed when running: %v", execError.Command))
		}
	} else {
		println("failed to down compose because compose is nil")
	}
}
