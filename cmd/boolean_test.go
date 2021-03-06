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

func TestBooleanCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "boolean",
			Output: "true\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "boolean -c 3",
			Output: "true\ntrue\nfalse\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "boolean --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
		{
			Name:   "-s flag",
			Input:  "boolean -s",
			Output: "T\n",
			IsErr:  false,
		},
		{
			Name:   "-s -c flags",
			Input:  "boolean -s -c 3",
			Output: "T\nT\nF\n",
			IsErr:  false,
		},
	}

	runTestTable(t, tests)
}

func TestCoinFlipCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "coin-flip",
			Output: "heads\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "coin-flip -c 3",
			Output: "heads\nheads\ntails\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "coin-flip --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
		{
			Name:   "-s flag",
			Input:  "coin-flip -s",
			Output: "H\n",
			IsErr:  false,
		},
		{
			Name:   "-s -c flags",
			Input:  "coin-flip -s -c 3",
			Output: "H\nH\nT\n",
			IsErr:  false,
		},
	}

	runTestTable(t, tests)
}

func TestGenderCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "gender",
			Output: "female\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "gender -c 3",
			Output: "female\nfemale\nmale\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "gender --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
		{
			Name:   "-s flag",
			Input:  "gender -s",
			Output: "F\n",
			IsErr:  false,
		},
		{
			Name:   "-s -c flags",
			Input:  "gender -s -c 3",
			Output: "F\nF\nM\n",
			IsErr:  false,
		},
	}

	runTestTable(t, tests)
}
