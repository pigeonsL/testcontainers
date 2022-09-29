package main

import (
	"fmt"
	"os"
	"testcontainers/cmd"
	"testcontainers/docker_compose"
)

func main() {
	cmd.Execute()
	var order, serviceName string
	for true {
		fmt.Println("There are the containers that have been started, Please choose whether you want to restart or exit the program：\n1.restart \n2.exit")
		fmt.Print("Please enter 1 or 2: ")
		fmt.Scanln(&order)
		if order == "1" {
			docker_compose.ShowDockerComposeService()
			fmt.Print("Please enter the name of the service you want to restart: ")
			fmt.Scanln(&serviceName)
			// todo check serviceName
			docker_compose.DockerComposeRestart(serviceName)
		} else if order == "2" {
			docker_compose.DockerComposeDown()
			os.Exit(1)
		} else {
			fmt.Println("Input error")
		}
	}
}
