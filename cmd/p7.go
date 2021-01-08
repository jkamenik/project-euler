package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p7Cmd = &cobra.Command{
	Use:   "p7",
	Short: "Problem 7",
	Long: `Problem 7

By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?`,

	RunE: func(cmd *cobra.Command, args []string) error {
		// maxPrime := 6
		maxPrime := 10_001
		countOfPrimes := 0
		lastPrime := 0

		for i := 2; i <= MaxInt; i++ {
			// Use the previously created for primefactors problem
			prime, _, _ := factors(i)

			if prime {
				log.Debug().Msgf("%d is prime", i)

				countOfPrimes = countOfPrimes + 1
				lastPrime = i

				if countOfPrimes >= maxPrime {
					break
				}
			}
		}

		log.Info().Msgf("The %d prime is %d", maxPrime, lastPrime)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(p7Cmd)
}
