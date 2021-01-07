package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"

	"github.com/rs/zerolog"
)

var rootCmd = &cobra.Command{
	Use:     "proect-euler",
	Version: "v0.0.1",
	Annotations: map[string]string{
		"container_repo": "jkamenik/project-euler",
	},
	Short: "Programming project euler as a set of commands",
	Long:  `Run individual Problems of Project Euler (https://projecteuler.net)`,

	PersistentPreRun: func(cmd *cobra.Command, arg []string) {
		if verbosity >= 1 {
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		}

		if verbosity >= 2 {
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		}

		log.Debug().Msgf("Verbosity: %d", verbosity)
	},
}

var verbosity = 0

// Execute runs the command
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().CountVarP(&verbosity, "verbose", "v", "More v's = more verbosity")
}
