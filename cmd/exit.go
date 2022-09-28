package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"testcontainers/docker_compose"
)

// exitCmd represents the exit command
var exitCmd = &cobra.Command{
	Use:   "exit",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		docker_compose.DockerComposeDown()
		fmt.Println("exit called")
	},
}

func init() {
	rootCmd.AddCommand(exitCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// exitCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// exitCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
