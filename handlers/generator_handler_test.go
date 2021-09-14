package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

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

			brackets := DefaultGenerate(tc.in)

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

func TestHandlerGenerator(t *testing.T) {
	t.Parallel()

	var testCases = []struct {
		name   string
		method string
		param  string
		gen    func(int) string
		body   string
		status int
	}{
		{"empty parameter", "GET", "", DefaultGenerate, "strconv.Atoi: parsing \"\": invalid syntax\n", 500},
		{"negative parameter", "GET", "-5", DefaultGenerate, "argument isn't positive: -5\n", 500},
		{"correct input", "GET", "4", func(l int) string {
			return "{])["
		}, "{])[", 200},
		{"unknown method", "POST", "2", DefaultGenerate, "method isn't allowed: POST\n", 405},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			req, err := http.NewRequest(tc.method, "/generate", nil)
			if err != nil {
				t.Fatal(err)
			}

			q := url.Values{}

			q.Add("n", tc.param)
			req.URL.RawQuery = q.Encode()

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
