// Copyright © 2017 Harry Bagdi <harry.bagdi@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/brianvoe/gofakeit"

	"github.com/spf13/cobra"
)

var short bool

// booleanCmd represents the boolean command
var booleanCmd = &cobra.Command{
	Use:   "boolean",
	Short: "Generate a random Boolean value.",
	Long: `Generates a random Boolean value.
Outcome will be either 'true' or 'false'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runBoolean, cmd, args)
	},
}

func runBoolean(cmd *cobra.Command, args []string) error {
	output := "T"
	val := gofakeit.Bool()
	if short {
		if !val {
			output = "F"
		}
		fmt.Fprintln(cmd.OutOrStdout(), output)
	} else {
		fmt.Fprintln(cmd.OutOrStdout(), val)
	}
	return nil
}

// coinFlipCmd represents the coin-flip command
var coinFlipCmd = &cobra.Command{
	Use:   "coin-flip",
	Short: "Generate an outcome of a coin flip.",
	Long: `Generates a random Boolean value.
Outcome will be either 'heads' or 'tails'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runCoinFlip, cmd, args)
	},
}

func runCoinFlip(cmd *cobra.Command, args []string) error {
	output := "H"
	val := gofakeit.Bool()
	if short {
		if !val {
			output = "T"
		}
		fmt.Fprintln(cmd.OutOrStdout(), output)
	} else {
		output = "heads"
		if !val {
			output = "tails"
		}
		fmt.Fprintln(cmd.OutOrStdout(), output)
	}
	return nil
}

// genderCmd represents the coin-flip command
var genderCmd = &cobra.Command{
	Use:   "gender",
	Short: "Generate a random gender(male/female only).",
	Long: `Generates a random Gender value.
Outcome will be either 'male' or 'female'.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runGender, cmd, args)
	},
}

func runGender(cmd *cobra.Command, args []string) error {
	output := "F"
	val := gofakeit.Bool()
	if short {
		if !val {
			output = "M"
		}
		fmt.Fprintln(cmd.OutOrStdout(), output)
	} else {
		output = "female"
		if !val {
			output = "male"
		}
		fmt.Fprintln(cmd.OutOrStdout(), output)
	}
	return nil
}

func init() {
	booleanCmd.Flags().BoolVarP(&short, "short", "s", false, "output an abbrevaited value")
	rootCmd.AddCommand(booleanCmd)

	coinFlipCmd.Flags().BoolVarP(&short, "short", "s", false, "output an abbrevaited value")
	rootCmd.AddCommand(coinFlipCmd)

	genderCmd.Flags().BoolVarP(&short, "short", "s", false, "output an abbrevaited value")
	rootCmd.AddCommand(genderCmd)
}
