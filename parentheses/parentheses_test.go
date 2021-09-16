package parentheses_test

import (
	"parentheses/parentheses"
	"strings"
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
		{"wrong with letters", "{(1 + 2) - [3 + 4}]", false},
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

func TestGenerateBrackets(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName string
		in       int
		out      int
	}{
		{"positive length", 8, 8},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			brackets := parentheses.GenerateBrackets(tc.in)

			if length := len(brackets); length != tc.in {
				t.Errorf("length want %v, length get %v", tc.in, length)
			}

			for _, char := range brackets {
				if !strings.ContainsRune("{}[]()", char) {
					t.Errorf("in sequence exists non bracket %v", char)
				}
			}
		})
	}
}

func TestGenerateBracketsPanic(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name        string
		length      int
		resultPanic bool
	}{
		{"panic zero", 0, true},
		{"panic negative", -1, true},
		{"don't panic", 1, false},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			defer func() {
				if wasPanic := recover() != nil; (!wasPanic || !tc.resultPanic) && (wasPanic || tc.resultPanic) {
					t.Errorf("panic want %v got %v", tc.resultPanic, wasPanic)
				}
			}()
			parentheses.GenerateBrackets(tc.length)
		})
	}
}
