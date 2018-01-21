package cmd

import (
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/stretchr/testify/assert"
)

func TestPasswordCmd(t *testing.T) {
	tests := []struct {
		Name      string
		Input     string
		Output    string
		SubString string
		IsErr     bool
	}{
		{
			Name:   "basic",
			Input:  "password",
			Output: "LE4C_4#032mcXPmR\n",
			IsErr:  false,
		},
		{
			Name:   "repeat",
			Input:  "password -c 2",
			Output: "LE4C_4#032mcXPmR\nhsS@nZXE?htbS-U&\n",
			IsErr:  false,
		},
		{
			Name:      "unknown flag",
			Input:     "password --unknown ",
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
