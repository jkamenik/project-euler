package cmd

import (
	"sort"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p3Cmd = &cobra.Command{
	Use:   "p3",
	Short: "Problem 3",
	Long: `Problem 3

	The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?`,

	RunE: func(cmd *cobra.Command, args []string) error {
		p, f1, f2 := factors(1)
		log.Debug().Msgf("Factors of 1: %d, %d (prime? %t)", f1, f2, p)

		p, f1, f2 = factors(2)
		log.Debug().Msgf("Factors of 2: %d, %d (prime? %t)", f1, f2, p)

		p, f1, f2 = factors(3)
		log.Debug().Msgf("Factors of 3: %d, %d (prime? %t)", f1, f2, p)

		p, f1, f2 = factors(13195)
		log.Debug().Msgf("Factors of 13195: %d, %d (prime? %t)", f1, f2, p)

		num := 600851475143
		factors := []int{}
		factors = primeFactors(factors, num)

		log.Debug().Msgf("Factors of %d: %v", num, factors)

		sort.Ints(factors)

		log.Info().Msgf("Largest prime factor of %d: %d", num, factors[len(factors)-1])

		return nil
	},
}

func primeFactors(seed []int, start int) []int {
	log.Debug().Msgf("prime factors of %d", start)
	log.Trace().Msgf("Factors found: %v", seed)

	prime, f1, f2 := factors(start)
	log.Debug().Msgf("Factors of %d: %d, %d (prime? %t)", start, f1, f2, prime)

	if prime {
		log.Debug().Msgf("Number is prime, inserting.")
		seed = append(seed, start)
		return seed
	}

	seed = primeFactors(seed, f1)
	seed = primeFactors(seed, f2)
	return seed
}

func factors(input int) (prime bool, fact1, fact2 int) {
	if input <= 3 {
		return true, 1, input
	}

	// this might be too much
	for fact1 = 2; fact1 < input; fact1++ {
		log.Trace().Msgf("Checking if %d is a factor of %d", fact1, input)

		if input%fact1 == 0 {
			fact2 = input / fact1
			prime = false

			log.Trace().Msgf("%d has the factors %d and %d", input, fact1, fact2)

			return
		}
	}

	return true, 1, input
}

func init() {
	rootCmd.AddCommand(p3Cmd)
}
