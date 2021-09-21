package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

var (
	errNotPositiveArgument  = errors.New("argument isn't positive")
	errMethodNotAllowed     = errors.New("method isn't allowed")
	errParameterNotProvided = errors.New("parameter isn't provided")
	errNotConvertableToInt  = errors.New("can't convert to int")
)

// GeneratorHandler returns body with the bracket sequence
// generated by Generate function.
type GeneratorHandler struct {
	gen Generator
}

// NewGenHandler is constructor for GeneratorHandler.
func NewGenHandler(gen Generator) *GeneratorHandler {
	return &GeneratorHandler{gen: gen}
}

// Generator is interface to generate string.
type Generator interface {
	Generate(length int) string
}

// ServeHTTP handles only GET method
// Writes sequence of brackets in body or error if it occurs.
func (g *GeneratorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	length, err := validateRequest(r)

	if errors.Is(err, errMethodNotAllowed) {
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)

		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	temp := g.gen.Generate(length)

	w.Write([]byte(temp)) //nolint:errcheck // error can't appear.
}

// function validateRequest checks if the method and parameter is correct.
// If the parameter "n" exists and if it is greater than 0.
// If error occur it returns (0, error), else it returns (value of parameter "n", nil).
func validateRequest(r *http.Request) (int, error) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return 0, fmt.Errorf("%w: %s", errMethodNotAllowed, r.Method)
	}

	key, exist := r.URL.Query()["n"]

	if !exist {
		return 0, fmt.Errorf("%w: %s", errParameterNotProvided, "n")
	}

	length, err := strconv.Atoi(key[0])

	if err != nil {
		return 0, fmt.Errorf("%w: %q", errNotConvertableToInt, key[0])
	}

	if length < 1 {
		return 0, fmt.Errorf("%w: %v", errNotPositiveArgument, length)
	}

	return length, nil
}
