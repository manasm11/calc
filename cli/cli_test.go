package cli

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		// test cases for add functionality
		{"./calc add 1 13", "14"},
		{"./calc add 22 8", "30"},
		{"./calc add a 8", "invalid integer 'a'" + usage},
		{"./calc add 1 *", "invalid integer '*'" + usage},
		{"./calc add ~ *", "invalid integer '~'" + "\ninvalid integer '*'" + usage},
		{"./calc add 1 5 6", "12"},

		// test cases for mul functionality
		{"./calc mul 1 13", "13"},
		{"./calc mul 22 8", "176"},
		{"./calc mul a 8", "invalid integer 'a'" + usage},
		{"./calc mul 1 *", "invalid integer '*'" + usage},
		{"./calc mul ~ *", "invalid integer '~'" + "\ninvalid integer '*'" + usage},
		{"./calc mul 2 5 6", "60"},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			s := strings.Split(c.input, " ")
			cli := NewCli()
			cli.Run(s)

			compareStrings(t, cli.output.String(), c.expected)
		})

		// if cli.output.String() != c.expected {
		// t.Errorf("Expected %s, got %s", c.expected, cli.output.String())
		// }
	}
}

// string comparison code from claude.ai
func compareStrings(t testing.TB, expected, actual string) {
	t.Helper()
	if expected == actual {
		return
	}

	// Split strings into lines for comparison
	expectedLines := strings.Split(expected, "\n")
	actualLines := strings.Split(actual, "\n")

	// Find the first different character and its position
	var diffIndex int
	var diffLine int
	var shortestLine string

	for i := 0; i < len(expectedLines) && i < len(actualLines); i++ {
		if expectedLines[i] != actualLines[i] {
			diffLine = i
			// Find the exact character where they differ
			minLen := len(expectedLines[i])
			if len(actualLines[i]) < minLen {
				minLen = len(actualLines[i])
				shortestLine = "actual"
			} else {
				shortestLine = "expected"
			}

			for j := 0; j < minLen; j++ {
				if expectedLines[i][j] != actualLines[i][j] {
					diffIndex = j
					break
				}
			}
			if diffIndex == 0 && len(expectedLines[i]) != len(actualLines[i]) {
				diffIndex = minLen
			}
			break
		}
	}

	// Print the difference with context
	t.Errorf("\nString difference found at line %d, position %d:\n", diffLine+1, diffIndex+1)

	// Print the expected line with marker
	if diffLine < len(expectedLines) {
		t.Errorf("Expected: %s\n%s↑\n", expectedLines[diffLine], strings.Repeat(" ", diffIndex))
	}

	// Print the actual line with marker
	if diffLine < len(actualLines) {
		t.Errorf("Actual:   %s\n%s↑\n", actualLines[diffLine], strings.Repeat(" ", diffIndex))
	}

	// Additional context for length differences
	if shortestLine != "" {
		t.Errorf("Note: %s line is shorter\n", shortestLine)
	}

	// Handle case where one string has more lines
	if len(expectedLines) != len(actualLines) {
		t.Errorf("\nLine count difference - Expected: %d, Actual: %d\n",
			len(expectedLines), len(actualLines))
	}
}
