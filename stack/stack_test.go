package stack_test

import (
	"errors"
	"parentheses/stack"
	"testing"
)

//func TestStack_Push(t *testing.T) {
//	t.Parallel()
//
//	var testCases = []struct{
//		testName string
//		in          []int
//		inCapacity  int
//	}
//}

func TestStack_Pop(t *testing.T) {
	t.Parallel()

	var testCases  = []struct{
		testName string
		pushed   []int
		len      int
		outVal   int
		outErr   error
	}{
		{"delete one element", []int{1,2,3}, 2, 3, nil},
		{"delete empty array one time", []int{},  0,0, stack.ErrPopEmpty},
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
			if !errors.Is(err,tc.outErr) {
				t.Errorf("error want %v, error get %v", tc.outErr, err)
			}
		})
	}
}

func TestStack_Empty(t *testing.T) {
	t.Parallel()

	var testCases = []struct{
		testName string
		in       []int
		out      bool
	}{
		{"empty", nil, true},
		{"elements exists", []int{1,2,3}, false},
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