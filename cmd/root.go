package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"testcontainers/docker_compose"
)

var rootCmd = &cobra.Command{
	Use:   "start",
	Short: "This will start a set of containers",
	Long:  "This is a bunch of containers about ROS, The three containers are Master, talker and listener.",
	Run: func(cmd *cobra.Command, args []string) {
		docker_compose.DockerComposeStart()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
