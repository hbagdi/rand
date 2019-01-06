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
	"bufio"
	"bytes"
	"os"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"

	"github.com/stretchr/testify/assert"
)

var (
	testOutBuffer bytes.Buffer
	testOut       *bufio.Writer
)

type testItem struct {
	Name        string
	Input       string
	Output      string
	SubString   string
	OutputRegex string
	IsErr       bool
}

func TestMain(m *testing.M) {
	testOut = bufio.NewWriter(&testOutBuffer)
	rootCmd.SetOutput(testOut)
	os.Exit(m.Run())
}

// testExecute is same as Execute() for testing purposes.
// It writes output into a buffer rather than stdout
// and also returns error (if any) from rootCmd
func testExecute(args string) (string, error) {
	count = 1
	testOutBuffer.Reset()
	rootCmd.SetArgs(strings.Split(args, " "))
	err := rootCmd.Execute()
	testOut.Flush()
	return testOutBuffer.String(), err
}

func runTestTable(t *testing.T, tests []testItem) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			a := assert.New(t)
			gofakeit.Seed(42)
			out, err := testExecute(test.Input)
			if test.OutputRegex != "" {
				a.Regexp(test.OutputRegex, out)
			}
			if test.Output != "" {
				a.Equal(test.Output, out)
			}
			if test.SubString != "" {
				a.Contains(out, test.SubString)
			}
			if test.IsErr {
				a.NotNil(err)
			} else {
				a.Nil(err)
			}
			resetGlobals()
		})
	}
}

func resetGlobals() {
	short = false
	count = 1
	hex = false
	rgb = false
	v6 = false
	paragraphs = 1
	sentences = 1
	words = 42
	template = ""
	templateFile = ""
	minLat = -90
	maxLat = 90
	minLong = -180
	maxLong = 180
}
