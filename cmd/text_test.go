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

func TestTextCmd(t *testing.T) {
	tests := []testItem{
		{
			Name:  "basic",
			Input: "text",
			Output: "At illum ut est sit soluta nulla numquam nobis " +
				"sunt quaerat ea dolores facere deleniti culpa numquam ut " +
				"distinctio maxime consequatur est qui corporis sunt officia " +
				"odit et quia odit molestias voluptas porro repellendus magnam " +
				"ipsa corporis eos rem non hic esse.\n",
			IsErr: false,
		},
		{
			Name:  "repeat",
			Input: "text -c 2",
			Output: "At illum ut est sit soluta nulla numquam nobis " +
				"sunt quaerat ea dolores facere deleniti culpa numquam ut " +
				"distinctio maxime consequatur est qui corporis sunt officia " +
				"odit et quia odit molestias voluptas porro repellendus magnam " +
				"ipsa corporis eos rem non hic esse.\n" +
				"Optio quisquam hic natus earum molestias iste architecto porro et " +
				"blanditiis iste eum repellendus nostrum qui eius suscipit fugit quia " +
				"quo et nesciunt quod fuga ut vel pariatur libero sequi rerum omnis " +
				"soluta facilis voluptatem possimus et voluptas eaque possimus harum " +
				"voluptatibus.\n",
			IsErr: false,
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
