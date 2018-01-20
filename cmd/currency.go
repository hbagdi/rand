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

var info bool

// currencyCmd represents the currency command
var currencyCmd = &cobra.Command{
	Use:   "currency",
	Short: "Generate a random currency",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runCurrency, cmd, args)
	},
}

func runCurrency(cmd *cobra.Command, args []string) error {
	currency := gofakeit.Currency()
	output := currency.Short
	if info {
		output += " " + currency.Long
	}
	fmt.Fprintln(cmd.OutOrStdout(), output)
	return nil
}

func init() {
	currencyCmd.Flags().BoolVarP(&info, "long", "l", false, "Add unabbreviated Currency name to output")
	rootCmd.AddCommand(currencyCmd)
}
