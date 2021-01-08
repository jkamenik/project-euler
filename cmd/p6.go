package cmd

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p6Cmd = &cobra.Command{
	Use:   "p6",
	Short: "Problem 6",
	Long: `Problem 6

	The sum of the squares of the first ten natural numbers is,
	
		1^2+2^2+...+10^2 = 385

	The square of the sum of the first ten natural numbers is,
	
		(1+2+...+10)^2 = 55^2 = 3025

	Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is

		3025 - 385 = 2640

	Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.`,

	RunE: func(cmd *cobra.Command, args []string) error {
		// max := 10
		max := 100

		sums := sumOfSquares(max)
		squares := SquareOfSums(max)

		log.Info().Msgf("Difference between sum of squares (%d) and square of sums (%d) for %d is %d", sums, squares, max, squares-sums)

		return nil
	},
}

func sumOfSquares(max int) int {
	acc := 0

	for i := 0; i <= max; i++ {
		acc = acc + (i * i)
	}

	return acc
}

func SquareOfSums(max int) int {
	acc := 0

	for i := 0; i <= max; i++ {
		acc = acc + i
	}

	return acc * acc
}

func init() {
	rootCmd.AddCommand(p6Cmd)
}
