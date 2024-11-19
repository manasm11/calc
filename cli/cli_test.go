package cli

import (
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"./calc add 1 13", "14"},
	}

	for _, c := range cases {
		s := strings.Split(c.input, " ")
		cli := NewCli()
		cli.Run(s)

		if cli.output.String() != c.expected {
			t.Errorf("Expected %s, got %s", c.expected, cli.output.String())
		}
	}
}
