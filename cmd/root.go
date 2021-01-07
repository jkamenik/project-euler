package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "proect-euler",
	Version: "v0.0.1",
	Annotations: map[string]string{
		"container_repo": "jkamenik/project-euler",
	},
	Short: "Programming project euler as a set of commands",
	Long:  `Run individual Problems of Project Euler (https://projecteuler.net)`,
}

// Execute runs the command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	// Global commands here
}
