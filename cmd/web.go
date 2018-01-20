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
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runIP, cmd, args)
	},
}

func runIP(cmd *cobra.Command, args []string) error {
	if v6 {
		fmt.Fprintln(cmd.OutOrStdout(), gofakeit.IPv6Address())
	} else {
		fmt.Fprintln(cmd.OutOrStdout(), gofakeit.IPv4Address())
	}
	return nil
}

var suffixOnly bool

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "Generate random Domain Name",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runDomain, cmd, args)
	},
}

func runDomain(cmd *cobra.Command, args []string) error {
	if suffixOnly {
		fmt.Fprintln(cmd.OutOrStdout(), gofakeit.DomainSuffix())
	} else {
		fmt.Fprintln(cmd.OutOrStdout(), gofakeit.DomainName())
	}
	return nil
}

// httpStatusCmd represents the http-status command
var httpStatusCmd = &cobra.Command{
	Use:   "http-status",
	Short: "Generate random HTTP Status Code",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runHTTPStatus, cmd, args)
	},
}

func runHTTPStatus(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(cmd.OutOrStdout(), gofakeit.StatusCode())
	return nil
}

// httpMethodCmd represents the http-method command
var httpMethodCmd = &cobra.Command{
	Use:   "http-method",
	Short: "Generate random HTTP Method",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runHTTPMethod, cmd, args)
	},
}

func runHTTPMethod(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(cmd.OutOrStdout(), gofakeit.HTTPMethod())
	return nil
}

// userAgentCmd represents the user-agent command
var userAgentCmd = &cobra.Command{
	Use:   "user-agent",
	Short: "Generate random User Agent",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runUserAgent, cmd, args)
	},
}

func runUserAgent(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(cmd.OutOrStdout(), gofakeit.UserAgent())
	return nil
}

func init() {
	ipCmd.Flags().BoolVar(&v6, "v6", false, "Output a IPv6 addresses")
	webCmd.AddCommand(ipCmd)

	domainCmd.Flags().BoolVar(&suffixOnly, "suffix-only", false, "Output only a TLD name")
	webCmd.AddCommand(domainCmd)

	webCmd.AddCommand(httpStatusCmd)

	webCmd.AddCommand(httpMethodCmd)

	webCmd.AddCommand(userAgentCmd)

	rootCmd.AddCommand(webCmd)
}
