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

func TestTemplateCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:      "empty",
			Input:     "template",
			SubString: "Error: Please provide a template via --template or --file",
			IsErr:     true,
		},
		{
			Name:   "basic",
			Input:  `template -t {person.first}###??`,
			Output: "Jeromy853pt\n",
			IsErr:  false,
		},
		{
			Name:      "file",
			Input:     `template -f nofile.txt`,
			SubString: "Error: open nofile.txt: no such file or directory\n",
			IsErr:     true,
		},
		{
			Name:   "file",
			Input:  `template -f testData/template.txt`,
			Output: "Greenfelder\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "text --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}
