package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p1Cmd = &cobra.Command{
	Use:   "p1",
	Short: "Problem 1",
	Long: `Problem 1

If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		factors := make([]int, 0)

		for i := 1; i < 1000; i++ {
			if i%3 == 0 {
				log.Debug().Msgf("%d is a factor of 3", i)
				factors = append(factors, i)

				// Don't double count
				continue
			}

			if i%5 == 0 {
				log.Debug().Msgf("%d is a factor of 5", i)
				factors = append(factors, i)
			}
		}
		log.Debug().Msgf("Factors: %v", factors)

		sum := 0

		for _, value := range factors {
			sum = sum + value
		}

		log.Info().Msgf("Answer: %d", sum)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(p1Cmd)
}
