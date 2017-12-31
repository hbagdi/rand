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
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"text/tabwriter"

	"os"

	"github.com/brianvoe/gofakeit"
	"github.com/spf13/cobra"
)

var template string
var templateFile string
var catagoriesHelp string

func initCatagoriesHelp() {
	var p string
	catagories := gofakeit.Catagories()
	var keys []string
	for k := range catagories {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		subCatagories := catagories[key]
		sort.Strings(subCatagories)
		p += key + ":\t" + strings.Join(subCatagories, ", ") + "\n"
	}

	var buf bytes.Buffer
	w := tabwriter.NewWriter(&buf, 0, 8, 0, '\t', 0)
	w.Write([]byte(p))
	w.Flush()
	catagoriesHelp = buf.String()
}

func init() {
	initCatagoriesHelp()
	var templateCmd = &cobra.Command{
		Use:   "template",
		Short: "Generate a random string based on a template",
		Long: `Generate a random string based on a template
A template is a simply a string with specific locations at which 
to insert random data. Place a catagory and sub-catagory in curly
braces {} and that part of the input will be replaced
Example: Input : {person.first}##??
		 Output: John42fs
Any # will be replaced by a numberic digit (0-9) 
and ? by alphabets (a-z).
Possible catagories and sub-catagories are:

` + catagoriesHelp,
		Run: func(cmd *cobra.Command, args []string) {
			if template == "" && templateFile == "" {
				cmd.Help()
				return
			}
			if template != "" {
				for i := 0; i < count; i++ {
					fmt.Println(gofakeit.Generate(template))
				}
				return
			}
			if templateFile != "" {
				content, err := ioutil.ReadFile(templateFile)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v", err)
					return
				}
				for i := 0; i < count; i++ {
					fmt.Println(gofakeit.Generate(string(content)))
				}
				return
			}
		},
	}

	templateCmd.Flags().StringVarP(&template,
		"template", "t", "", "Generate a random string based on a template")
	templateCmd.Flags().StringVarP(&templateFile, "template-file", "f", "",
		"Generate a random string based on a template from a file")
	rootCmd.AddCommand(templateCmd)
}
