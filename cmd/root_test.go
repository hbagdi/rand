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

	"github.com/stretchr/testify/assert"
)

var (
	testOutBuffer bytes.Buffer
	testOut       *bufio.Writer
)

func TestMain(m *testing.M) {
	testOut = bufio.NewWriter(&testOutBuffer)
	rootCmd.SetOutput(testOut)
	os.Exit(m.Run())
}

// testExecute is same as Execute() for testing purposes.
// It writes output into a buffer rather than stdout
// and also returns error (if any) from rootCmd
func testExecute(args string) (string, error) {
	testOutBuffer.Reset()
	rootCmd.SetArgs(strings.Split(args, " "))
	err := rootCmd.Execute()
	testOut.Flush()
	return testOutBuffer.String(), err
}

func TestRootCmd(t *testing.T) {
	tests := []struct {
		Name      string
		Input     string
		Output    string
		SubString string
		IsErr     bool
	}{
		{
			Name:      "basic",
			Input:     "",
			SubString: rootCmd.Short,
			IsErr:     false,
		},
		{
			Name:   "-v flag",
			Input:  "-v",
			Output: VERSION + "\n",
			IsErr:  false,
		},
		{
			Name:   "--version flag",
			Input:  "--version",
			Output: VERSION + "\n",
			IsErr:  false,
		},
		{
			Name:      "-h flag",
			Input:     "-h",
			SubString: rootCmd.Short,
			IsErr:     false,
		},
		{
			Name:      "--help flag",
			Input:     "--help",
			SubString: rootCmd.Short,
			IsErr:     false,
		},
		{
			Name:      "non-existing flag",
			Input:     "--random",
			SubString: "Error: unknown flag: --random",
			IsErr:     true,
		},
		{
			Name:      "repeat",
			Input:     "-c 2",
			SubString: rootCmd.Short,
			IsErr:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			a := assert.New(t)
			out, err := testExecute(test.Input)
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
		})
	}
}
