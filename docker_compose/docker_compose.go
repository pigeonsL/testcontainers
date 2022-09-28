package docker_compose

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/testcontainers/testcontainers-go"
	"strings"
)

type rosContainer struct {
	testcontainers.LocalDockerCompose
}

func DockerCompose_Start() *rosContainer {
	composeFilePaths := []string{"docker/docker-compose.yml"}
	identifier := strings.ToLower(uuid.New().String())

	compose := testcontainers.NewLocalDockerCompose(composeFilePaths, identifier)
	execError := compose.
		WithCommand([]string{"up", "-d"}).
		WithEnv(map[string]string{
			"key1": "value1",
			"key2": "value2",
		}).
		Invoke()
	err := execError.Error
	if err != nil {
		println(fmt.Errorf("could not run compose file: %v - %v", composeFilePaths, err))
		return nil
	}
	return &rosContainer{LocalDockerCompose: *compose}
}

func DockerCompose_Down(compose *rosContainer) {
	if compose != nil {
		execError := compose.Down()
		if execError.Error != nil {
			_ = fmt.Errorf("failed when running: %v", execError.Command)
		}
	}
}
