package cmd

import (
	"testing"
)

func TestConfirmation(t *testing.T) {
	for _, test := range []struct {
		label    string
		in       []string
		err      error
		expected string
	}{
		{
			"test default",
			[]string{"port", "host"},
			nil,
			"--http-disallow-url /_health --http-disallow-url /_metrics",
		},
		{
			"test 3 arguments (1 filter)",
			[]string{"port", "host", "/get_this"},
			nil,
			"--http-allow-url /get_this",
		},
		{
			"test 4 arguments (2 filters)",
			[]string{"port", "host", "/get_this", "/and_this"},
			nil,
			"--http-allow-url /get_this --http-allow-url /and_this",
		},
		{
			"test not enough arguments",
			[]string{"port"},
			ErrNotEnoughArgs,
			"",
		},
	} {
		config, err := createConfig(test.in)
		if config.Filter != test.expected {
			t.Fatalf("%s value check failed. Expected %s, got %s", test.label, test.expected, config.Filter)
		}
		if err != test.err {
			t.Fatalf("%s error check failed. Expected %s, got %s", test.label, test.err, err)
		}
	}
}
