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
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"text/tabwriter"

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
)

var template string
var templateFile string
var categoriesHelp string

func initCategoriesHelp() {
	var p string
	categories := gofakeit.Categories()
	var keys []string
	for k := range categories {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		subCategories := categories[key]
		sort.Strings(subCategories)
		p += key + ":\t" + strings.Join(subCategories, ", ") + "\n"
	}

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 8, 0, '\t', 0)
	w.Write([]byte(p))
	w.Flush()
	categoriesHelp = buf.String()
}

func runTemplate(cmd *cobra.Command, args []string) error {
	if template == "" && templateFile == "" {
		return errors.New("Please provide a template" +
			" via --template or --file")
	}
	var content string
	if template != "" {
		content = template
	} else {
		fileContent, err := ioutil.ReadFile(templateFile)
		if err != nil {
			return err
		}
		content = string(fileContent)
	}
	return repeatFunc(
		func(cmd *cobra.Command, args []string) error {
			fmt.Fprintln(cmd.OutOrStdout(), gofakeit.Generate(string(content)))
			return nil
		}, cmd, args)
}

func init() {
	initCategoriesHelp()
	var templateCmd = &cobra.Command{
		Use:   "template",
		Short: "Generate a random string based on a template",
		Long: `Generate a random string based on a template
A template is a simply a string with specific locations at which 
to insert random data. Place a category and sub-category in curly
braces {} and that part of the input will be replaced
Example: Input : {person.first}##??
		 Output: John42fs
Any # will be replaced by a numberic digit (0-9) 
and ? by alphabets (a-z).
Possible categories and sub-categories are:

` + categoriesHelp,
		RunE: runTemplate,
	}

	templateCmd.Flags().StringVarP(&template,
		"template", "t", "", "Generate a random string based on a template")
	templateCmd.Flags().StringVarP(&templateFile, "file", "f", "",
		"Generate a random string based on a template from an input file")
	rootCmd.AddCommand(templateCmd)
}
