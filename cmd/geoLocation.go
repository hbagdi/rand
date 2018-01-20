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

var (
	minLat  = -90.0
	maxLat  = 90.0
	minLong = -180.0
	maxLong = 180.0
)

// geoLocationCmd represents the geo-location command
var geoLocationCmd = &cobra.Command{
	Use:   "geo-location",
	Short: "Generate a random Geographic coordinate location (lat,long)",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runGeoLoc, cmd, args)
	},
}

func runGeoLoc(cmd *cobra.Command, args []string) error {
	lat, err := gofakeit.LatitudeInRange(minLat, maxLat)
	if err != nil {
		return err
	}

	long, err := gofakeit.LongitudeInRange(minLong, maxLong)
	if err != nil {
		return err
	}

	latString := strconv.FormatFloat(lat, 'f', 6, 64)
	longString := strconv.FormatFloat(long, 'f', 6, 64)
	fmt.Fprintln(cmd.OutOrStdout(), "("+latString+","+longString+")")
	return nil
}

func init() {
	geoLocationCmd.Flags().Float64Var(&minLat, "min-lat", minLat, "minimum of latitude range")
	geoLocationCmd.Flags().Float64Var(&maxLat, "max-lat", maxLat, "maximum of latitude range")
	geoLocationCmd.Flags().Float64Var(&minLong, "min-long", minLong, "minimum of longitude range")
	geoLocationCmd.Flags().Float64Var(&maxLong, "max-long", maxLong, "maximum of longitude range")

	rootCmd.AddCommand(geoLocationCmd)
}
