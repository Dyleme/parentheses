package handlers

import (
	"errors"
	"net/http/httptest"
	"testing"
)

type mockGenerator string

func (f mockGenerator) Generate(length int) string {
	return string(f)
}

func TestHandlerGenerator(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name      string
		method    string
		parameter string
		value     string
		gen       Generator
		body      string
		status    int
	}{
		{"empty parameter", "GET", "n", "", &BracketsGenerator{}, "can't convert to int: \"\"\n", 500},
		{"negative parameter", "GET", "n", "-5", &BracketsGenerator{}, "argument isn't positive: -5\n", 500},
		{"correct input", "GET", "n", "4", mockGenerator("{])["), "{])[", 200},
		{"unknown method", "POST", "n", "2", &BracketsGenerator{}, "method isn't allowed: POST\n", 405},
		{"wrong parameter", "GET", "t", "2", &BracketsGenerator{}, "parameter isn't provided: n\n", 500},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			path := "/generate?" + tc.parameter + "=" + tc.value
			req := httptest.NewRequest(tc.method, path, nil)

			rr := httptest.NewRecorder()

			handler := GeneratorHandler{tc.gen}

			handler.ServeHTTP(rr, req)

			if rr.Code != tc.status {
				t.Errorf("status code want %v get %v", tc.status, rr.Code)
			}

			if rr.Body.String() != tc.body {
				t.Errorf("body want %q get %q", tc.body, rr.Body.String())
			}
		})
	}
}

func TestValidateRequest(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		testName  string
		method    string
		parameter string
		value     string
		err       error
		length    int
	}{
		{"correct", "GET", "n", "8", nil, 8},
		{"wrong method", "POST", "n", "12", errMethodNotAllowed, 0},
		{"parameter not exist", "GET", "t", "8", errParameterNotProvided, 0},
		{"empty parameter", "GET", "n", "", errNotConvertableToInt, 0},
		{"string parameter", "GET", "n", "string", errNotConvertableToInt, 0},
		{"negative parameter", "GET", "n", "-2", errNotPositiveArgument, 0},
		{"zero parameter", "GET", "n", "0", errNotPositiveArgument, 0},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			target := "/generate?" + tc.parameter + "=" + tc.value

			req := httptest.NewRequest(tc.method, target, nil)

			length, err := validateRequest(req)

			if !errors.Is(err, tc.err) {
				t.Errorf("want error %v get error %v", tc.err, err)
			}

			if length != tc.length {
				t.Errorf("want length %v get length %v", tc.length, length)
			}
		})
	}
}
