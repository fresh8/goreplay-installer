package cmd

import (
	"testing"
)

func TestConfirmation(t *testing.T) {
	for _, test := range []struct {
		label string
		in    []string
		out   string
	}{
		{
			"test default",
			[]string{"port", "host"},
			"--http-disallow-url /_health --http-disallow-url /_metrics",
		},
		{
			"test 1 param",
			[]string{"port", "host", "/get_this"},
			"--http-allow-url /get_this",
		},
		{
			"test 2 params",
			[]string{"port", "host", "/get_this", "/and_this"},
			"--http-allow-url /get_this --http-allow-url /and_this",
		},
	} {
		config := createConfig(test.in)
		if config.Filter != test.out {
			t.Fatalf("%s failed. Expected %s, got %s", test.label, test.out, config.Filter)
		}
	}
}
