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

func TestWebCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:      "basic",
			Input:     "web",
			SubString: webCmd.Short,
			IsErr:     false,
		},
		{
			Name:      "repeat",
			Input:     "web -c 2",
			SubString: webCmd.Short,
			IsErr:     false,
		},
		{
			Name:      "unknown flag",
			Input:     "geo-location --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}

func TestIPCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "web ip",
			Output: "249.195.38.38\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "web ip -c 2",
			Output: "249.195.38.38\n45.21.213.246\n",
			IsErr:  false,
		},
		{
			Name:   "-v6 flag",
			Input:  "web ip --v6",
			Output: "2001:cafe:64b1:b44b:f284:923e:d7df:7a61\n",
			IsErr:  false,
		},
		{
			Name:  "-v6 -c flags",
			Input: "web ip -c 2 --v6",
			Output: "2001:cafe:64b1:b44b:f284:923e:d7df:7a61\n" +
				"2001:cafe:5ba5:8588:9b70:75d3:19f9:826f\n",
			IsErr: false,
		},
		{
			Name:      "unknown flag",
			Input:     "web ip --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}

func TestDomainCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "web domain",
			Output: "humanmagnetic.net\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "web domain -c 2",
			Output: "humanmagnetic.net\nlegacytransform.com\n",
			IsErr:  false,
		},
		{
			Name:   "--suffix-only flag",
			Input:  "web domain --suffix-only",
			Output: "org\n",
			IsErr:  false,
		},
		{
			Name:   "--suffix-only -c flags",
			Input:  "web domain -c 2 --suffix-only",
			Output: "org\ninfo\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "web domain --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}

func TestHTTPMethodCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "web http-method",
			Output: "DELETE\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "web http-method -c 3",
			Output: "DELETE\nDELETE\nPOST\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "web http-method --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}

func TestHTTPStatusCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:   "basic",
			Input:  "web http-status",
			Output: "205\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "web http-status -c 3",
			Output: "205\n201\n403\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "web http-status --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}

func TestUserAgentCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:  "basic",
			Input: "web user-agent",
			Output: "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) " +
				"AppleWebKit/5362 (KHTML, like Gecko) " +
				"Chrome/36.0.823.0 Mobile Safari/5362\n",
			IsErr: false,
		},
		{
			Name:  "repeat",
			Input: "web user-agent -c 2",
			Output: "Mozilla/5.0 (Macintosh; PPC Mac OS X 10_6_5) " +
				"AppleWebKit/5362 (KHTML, like Gecko) " +
				"Chrome/36.0.823.0 Mobile Safari/5362\n" +
				"Mozilla/5.0 (Macintosh; U; PPC Mac OS X 10_9_9 rv:6.0; en-US) " +
				"AppleWebKit/534.44.5 (KHTML, like Gecko) " +
				"Version/4.2 Safari/534.44.5\n",
			IsErr: false,
		},
		{
			Name:      "unknown flag",
			Input:     "web user-agent --unknown ",
			SubString: "Error: unknown flag: --unknown",
			IsErr:     true,
		},
	}

	runTestTable(t, tests)
}
