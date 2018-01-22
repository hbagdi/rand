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

func TestCreditCardCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "credit-card",
			Output: "MasterCard 2720995347682712 508 04/22\n",
			IsErr:  false,
		},
		{
			Name:  "repeat",
			Input: "credit-card -c 2",
			Output: "MasterCard 2720995347682712 508 04/22\n" +
				"MasterCard 6527610785142054 401 07/23\n",
			IsErr: false,
		},
		{
			Name:      "unknown flag",
			Input:     "credit-card --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}
