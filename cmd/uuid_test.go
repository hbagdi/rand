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
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestUUIDCmd(t *testing.T) {
	tests := []struct {
		Name      string
		Input     string
		Output    string
		SubString string
		IsErr     bool
	}{
		{
			Name:   "basic",
			Input:  "uuid",
			Output: "538c7f96-b164-4f1b-97bb-9f4bb472e89f\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "uuid -c 2",
			Output: "538c7f96-b164-4f1b-97bb-9f4bb472e89f\n5b1484f2-5209-49d9-b43e-92ba09dd9d52\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "uuid --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			a := assert.New(t)
			gofakeit.Seed(42)
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
