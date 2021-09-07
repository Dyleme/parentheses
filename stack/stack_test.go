package stack_test

import (
	"errors"
	"parentheses/stack"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName string
		capacity int
	}{
		{"zero capacity", 0},
		{"capacity", 13},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			s := stack.New(tc.capacity)

			if result := cap(*s); result != tc.capacity {
				t.Errorf("want cap %v, get cap %v", tc.capacity, result)
			}

			if result := len(*s); result != 0 {
				t.Errorf("want len %v, get len %v", 0, result)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func TestStack_Push(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName string
		pushed   []int
		capacity int
		result   []int
	}{
		{"push one element", []int{1}, 3, []int{1}},
		{"push slice", []int{11, 22, 33, 44}, 5, []int{11, 22, 33, 44}},
		{"relocate slice", []int{1, 2, 3, 4, 5, 6}, 3, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			s := stack.New(tc.capacity)

			for _, element := range tc.pushed {
				s.Push(element)
			}

			if !equal(*s, tc.result) {
				t.Errorf("stack: want result %v, get result %v", tc.result, *s)
			}
		})
	}
}

func TestStack_Top(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName string
		inSlice  []int
		outVal   int
		outErr   error
	}{
		{"empty", nil, 0, stack.ErrEmptyStack},
		{"elements exists", []int{1, 2, 3}, 3, nil},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			var s = stack.NewFromSlice(tc.inSlice)

			result, err := s.Top()

			if result != tc.outVal {
				t.Errorf("stack: want value %v, get value %v", tc.outVal, result)
			}

			if !errors.Is(err, tc.outErr) {
				t.Errorf("stack: want error %v, get error %v", tc.outErr, err)
			}
		})
	}
}

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName string
		pushed   []int
		len      int
		outVal   int
		outErr   error
	}{
		{"delete one element", []int{1, 2, 3}, 2, 3, nil},
		{"delete empty array one time", []int{}, 0, 0, stack.ErrEmptyStack},
		{"delete one element", []int{1}, 0, 1, nil},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			var s = stack.NewFromSlice(tc.pushed)

			var val int
			var err error
			val, err = s.Pop()
			if val != tc.outVal {
				t.Errorf("value want %v, value get %v", tc.outVal, val)
			}
			if !errors.Is(err, tc.outErr) {
				t.Errorf("error want %v, error get %v", tc.outErr, err)
			}
		})
	}
}

func TestStack_Empty(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName string
		in       []int
		out      bool
	}{
		{"empty", nil, true},
		{"elements exists", []int{1, 2, 3}, false},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			var s = stack.NewFromSlice(tc.in)

			if result := s.Empty(); result != tc.out {
				t.Errorf("stack: want result %v, get result %v", tc.out, result)
			}
		})
	}
}
