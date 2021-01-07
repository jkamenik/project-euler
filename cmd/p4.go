package cmd

import (
	"fmt"
	"sort"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var p4Cmd = &cobra.Command{
	Use:   "p4",
	Short: "Problem 4",
	Long: `Problem 4

A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
`,

	RunE: func(cmd *cobra.Command, args []string) error {
		isPalndromicInt(1)
		isPalndromicInt(101)
		isPalndromicInt(102)

		max := 999
		prod := []int{}

		for i := max; i >= 1; i-- {
			for j := max; j >= 1; j-- {
				product := i * j
				log.Trace().Msgf("Product of %d, and %d is %d", i, j, product)

				if isPalndromicInt(product) {
					log.Debug().Msgf("%d is a palindrome", product)
					prod = append(prod, product)
					// return nil
				}
			}
		}

		sort.Ints(prod)

		log.Info().Msgf("Largest palindrome: %d", prod[len(prod)-1])

		return nil
	},
}

func isPalndromicInt(num int) bool {
	s := fmt.Sprintf("%d", num)

	runes := []rune(s)
	r := []rune{}
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, runes[i])
	}

	palendromic := s == string(r)

	log.Debug().Msgf("Palendromic? %t, reverse for %s is %v", palendromic, s, string(r))

	return palendromic
}

func init() {
	rootCmd.AddCommand(p4Cmd)
}
