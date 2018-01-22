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

func TestETimeCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "time",
			Output: "1905-04-24 06:43:25.745640357 +0000 UTC\n",
			IsErr:  false,
		},
		{
			Name:  "repeat",
			Input: "time -c 2",
			Output: "1905-04-24 06:43:25.745640357 +0000 UTC\n" +
				"2016-04-15 01:47:15.260356828 +0000 UTC\n",
			IsErr: false,
		},
		{
			Name:      "unknown flag",
			Input:     "time --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}
