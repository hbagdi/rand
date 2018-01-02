// Copyright © 2018 Harry Bagdi <harry.bagdi@gmail.com>
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
	"math"

	"github.com/brianvoe/gofakeit"

	"github.com/spf13/cobra"
)

var min = math.MinInt32
var max = math.MaxInt32

// numberCmd represents the number command
var numberCmd = &cobra.Command{
	Use:   "number",
	Short: "Generate a random integer",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			fmt.Println(gofakeit.Number(min, max))
		}
	},
}

func init() {
	numberCmd.Flags().IntVar(&min, "min", min, "minimum of the range")
	numberCmd.Flags().IntVar(&max, "max", max, "maximum of the range")
	rootCmd.AddCommand(numberCmd)
}
