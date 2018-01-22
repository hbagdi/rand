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

var words int
var sentences int
var paragraphs int

// textCmd represents the text command
var textCmd = &cobra.Command{
	Use:   "text",
	Short: "Generate filler text based on Lorem Ipsum",
	RunE: func(cmd *cobra.Command, args []string) error {
		return repeatFunc(runText, cmd, args)
	},
}

func runText(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(cmd.OutOrStdout(), gofakeit.Paragraph(paragraphs,
		sentences, words, "\n\n"))
	return nil
}

func init() {
	textCmd.Flags().IntVarP(&words, "words", "w", 42,
		"Number of words per sentence")
	textCmd.Flags().IntVarP(&sentences, "sentences", "s", 1,
		"Number of senetences per paragraph")
	textCmd.Flags().IntVarP(&paragraphs, "paragraphs", "p", 1,
		"Number of paragraphs in the output")
	rootCmd.AddCommand(textCmd)
}
