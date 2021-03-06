package cmd_test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/trivago/hamara/cmd"
)

func TestRootCmd(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		fixture string
	}{
		{"no arguments", []string{}, "no-args.golden"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)

			actual := new(bytes.Buffer)
			rootCmd := cmd.NewRootCmd(tt.args)
			rootCmd.SetOutput(actual)
			rootCmd.SetArgs(tt.args)
			rootCmd.Execute()

			expected, err := ioutil.ReadFile("testdata/" + tt.fixture)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(string(expected), actual.String())
		})
	}

}
