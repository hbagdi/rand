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
	"strconv"

	"github.com/brianvoe/gofakeit"

	"github.com/spf13/cobra"
)

// geoLocationCmd represents the geo-location command
var geoLocationCmd = &cobra.Command{
	Use:   "geo-location",
	Short: "Generate a random Geographic coordinate location (lat,long)",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			lat := strconv.FormatFloat(gofakeit.Latitude(), 'f', 6, 64)
			long := strconv.FormatFloat(gofakeit.Longitude(), 'f', 6, 64)
			fmt.Println("(" + lat + "," + long + ")")
		}
	},
}

func init() {
	rootCmd.AddCommand(geoLocationCmd)
}
