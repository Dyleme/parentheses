package balanced_test

import (
	"testing"

	bb "github.com/Dyleme/parentheses/balanced"
)

func TestIsBalanced(t *testing.T) {
	var testCases = []struct {
		testName string
		testString string
		out bool
	}{
		{"standard", "{}", true},
		{"1", "[{()}]", true},
		{"2", "{[}]", false},
		{"3", "{{}}", true},
		{"4", "(){}", true},
		{"5", "({}(", false},
		{"6", "([[)", false},
		{"7", ")[])", false},
		{"8", ")", false},
		{"9", "(", false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			if result := bb.IsBalanced(tc.testString); result == tc.out {
				t.Errorf("balanced: want result %v, get result %v", tc.out, result)
			}
		})
	}
}