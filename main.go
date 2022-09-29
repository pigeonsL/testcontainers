package main

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
	"testcontainers/docker_compose"
)

func HomePageRun() (int, string, error) {
	prompt := promptui.Select{
		Label: "Select command",
		Items: []string{"restart", "exit"},
	}

	return prompt.Run()
}

func RestartPageRun() (int, string, error) {
	service := promptui.Select{
		Label: "Select you want restart service name",
		Items: []string{"master", "talker", "listener"},
	}
	return service.Run()
}

func main() {
	docker_compose.DockerComposeStart()
	for true {
		_, command, err := HomePageRun()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		if command == "restart" {
			_, serviceName, err := RestartPageRun()
			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}
			docker_compose.DockerComposeRestart(serviceName)
		}

		if command == "exit" {
			docker_compose.DockerComposeDown()
			fmt.Println("The program has exited.")
			os.Exit(1)
		}
	}
}
