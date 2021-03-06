package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p2Cmd = &cobra.Command{
	Use:   "p2",
	Short: "Problem 2",
	Long: `Problem 2

	Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:

		1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

	By considering the terms in the Fibonacci sequence whose values do not exceed four million, find the sum of the even-valued terms.
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		max := 4_000_000
		fibs := []int{1, 2} // need a seed

		for true {
			next := fibs[len(fibs)-1] + fibs[len(fibs)-2]

			log.Trace().Msgf("Next: %d", next)

			if next >= max {
				log.Debug().Msgf("%d is greater then max of %d", next, max)
				break
			}

			fibs = append(fibs, next)
		}

		log.Debug().Msgf("All Fibs below %d: %v", max, fibs)

		sum := 0

		for _, value := range fibs {
			if value%2 != 0 {
				log.Trace().Msgf("Skipping odd number %d", value)
				continue
			}

			sum = sum + value
			log.Debug().Msgf("Sum plus %d: %d", value, sum)
		}

		log.Info().Msgf("Answer: %d", sum)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(p2Cmd)
}
