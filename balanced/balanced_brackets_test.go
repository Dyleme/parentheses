package balanced

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName   string
		testString string
		out        bool
	}{
		{"{}", "{}", true},
		{"[{()}]", "[{()}]", true},
		{"{[}]", "{[}]", false},
		{"{{}}", "{{}}", true},
		{"(){}", "(){}", true},
		{"({}(", "({}(", false},
		{"([[}", "([[)", false},
		{")[])", ")[])", false},
		{")", ")", false},
		{"(", "(", false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			if result := IsBalanced(tc.testString); result != tc.out {
				t.Errorf("balanced: want result %v, get result %v", tc.out, result)
			}
		})
	}
}
