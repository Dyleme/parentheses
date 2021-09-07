package balanced_test

import (
	parentheses "parentheses/parentheses"

	"testing"
)

func TestIsBalanced(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName   string
		testString string
		out        bool
	}{
		{"all bracket are balanced at center", "[{()}]", true},
		{"wrong sequence of close brackets", "{[}]", false},
		{"two times correct brackets", "(){}", true},
		{"wrapped brackets are unclosed", "({}(", false},
		{"inner brackets are unclosed", "([[)", false},
		{"wrapped bracket are opened", ")[])", false},
		{"one close bracket", ")", false},
		{"one open bracket", "(", false},
		{"true with letters", "(1 + 2 * {3 + 4}) * [3 + 2]", true},
		{"wrong with letters", "{(1 + 2) - [3 + 4}]",false},
		{"UTF-8 characters", "(П{Р}[ивет])", true},
		{"empty string", "", true},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			if result := parentheses.IsBalanced(tc.testString); result != tc.out {
				t.Errorf("balanced: want result %v, get result %v", tc.out, result)
			}
		})
	}
}
