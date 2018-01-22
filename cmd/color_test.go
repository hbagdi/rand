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
)

func TestColorCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "color",
			Output: "Sienna\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "color -c 2",
			Output: "Sienna\nNavajoWhite\n",
			IsErr:  false,
		},
		{
			Name:   "--hex flag",
			Input:  "color --hex",
			Output: "#68pt27\n",
			IsErr:  false,
		},
		{
			Name:   "--rgb flag",
			Input:  "color --rgb",
			Output: "rgb(177,75,132)\n",
			IsErr:  false,
		},
		{
			Name:   "--hex -c flags",
			Input:  "color --hex -c 3",
			Output: "#68pt27\n#8sq2v7\n#j0k540\n",
			IsErr:  false,
		},
		{
			Name:   "--rgb -c flags",
			Input:  "color --rgb -c 3",
			Output: "rgb(177,75,132)\nrgb(62,223,97)\nrgb(165,136,112)\n",
			IsErr:  false,
		},
		{
			Name:      "--rgb --hex flags",
			Input:     "color --rgb --hex",
			SubString: "Error: --hex and --rgb are exclusive",
			IsErr:     true,
		},
		{
			Name:      "unknown flag",
			Input:     "color --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}
