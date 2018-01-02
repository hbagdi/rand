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

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Generate random values for all things web",
	Long:  ``,
}

var v6 bool

// ipCmd represents the ip command
var ipCmd = &cobra.Command{
	Use:   "ip",
	Short: "Generate random IP v4/v6 address",
	Long: `Generate a random IP Address
IP v4(default) and v6 are supported.
Use the following flag for IPv6`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			if v6 {
				fmt.Println(gofakeit.IPv6Address())
			} else {
				fmt.Println(gofakeit.IPv4Address())
			}
		}
	},
}

var suffixOnly bool

// domainCmd represents the ip command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Generate random Domain Name",
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < count; i++ {
			if suffixOnly {
				fmt.Println(gofakeit.DomainSuffix())
			} else {
				fmt.Println(gofakeit.DomainName())
			}
		}
	},
}

func init() {
	ipCmd.Flags().BoolVar(&v6, "v6", false, "Output a IPv6 addresses")
	webCmd.AddCommand(ipCmd)

	domainCmd.Flags().BoolVar(&suffixOnly, "suffix-only", false, "Output only a TLD name")
	webCmd.AddCommand(domainCmd)

	rootCmd.AddCommand(webCmd)
}
