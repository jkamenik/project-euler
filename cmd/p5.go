package cmd

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Get the max unsigned, then shift positive (leading 0) and cast to signed.
// This makes the max signed int
const MaxInt = int(^uint(0) >> 1)

var p5Cmd = &cobra.Command{
	Use:   "p5",
	Short: "Problem 5",
	Long: `Problem 5

	2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

	What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?`,

	RunE: func(cmd *cobra.Command, args []string) error {
		// maxDivisor := 2
		// maxDivisor := 10
		maxDivisor := 20

		// seed := 2
		seed := 2520

		num, err := smallestMultiple(maxDivisor, seed)
		if err != nil {
			return err
		}

		log.Info().Msgf("Smallest Multiple for all numbers 1 - %d is %d", maxDivisor, num)
		return nil
	},
}

func smallestMultiple(maxDivisor, seed int) (int, error) {
	log.Debug().Msgf("smallestMultiple(%d, %d)", maxDivisor, seed)
	d := true
	divides := &d

	for num := seed; num < MaxInt; num++ {
		*divides = true

		for i := 2; i <= maxDivisor; i++ {
			log.Trace().Msgf("Checking %d is a multiple of %d", i, num)
			if num%i != 0 {
				log.Debug().Msgf("%d is not a mulitple of %d, bailing", i, num)
				*divides = false
				break
			}
		}

		if *divides {
			log.Debug().Msgf("returning %d", num)
			return num, nil
		}
	}

	return 0, fmt.Errorf("No divisor found within the max Int (%d)", MaxInt)
}

func init() {
	rootCmd.AddCommand(p5Cmd)
}
