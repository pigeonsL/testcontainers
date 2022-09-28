package main

import (
	"fmt"
	"os"
	"testcontainers/cmd"
	"testcontainers/docker_compose"
)

func main() {
	cmd.Execute()
	var order, containerId string
	for true {
		fmt.Println("There are the containers that have been started, Please choose whether you want to restart or exit the program：\n restart service \n exit")
		fmt.Scanln(&order, &containerId)
		if order == "restart" {
			docker_compose.DockerComposeRestart(containerId)
		} else if order == "exit" {
			docker_compose.DockerComposeDown()
			os.Exit(1)
		} else {
			fmt.Println("Input error")
		}
	}
}
