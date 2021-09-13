package handlers

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var (
	ErrNotPositiveArgument   = errors.New("argument isn't positive")
	ErrMethodIsntAllowed     = errors.New("method isn't allowed")
	ErrParameterIsntProvided = errors.New("parameter isn't provided")
)

type GeneratorHandler struct {
	Generate func(length int) string
}

func DefaultGenerate(length int) string {
	var bracket = "(){}[]"

	rand.Seed(time.Now().Unix())

	var sb strings.Builder

	sb.Grow(length)

	for i := 0; i < length; i++ {
		sb.WriteRune(rune(bracket[rand.Intn(len(bracket))]))
	}

	return sb.String()
}

// ServeHTTP handles only GET method
// Writes sequence of brackets.
func (g *GeneratorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	length, err := validateRequest(r)

	switch {
	case errors.Is(err, ErrMethodIsntAllowed):
		http.Error(w, err.Error(), http.StatusMethodNotAllowed)
	case err != nil:
		http.Error(w, err.Error(), http.StatusInternalServerError)
	default:
		temp := g.Generate(length)
		_, errWrite := w.Write([]byte(temp))

		if errWrite != nil {
			http.Error(w, errWrite.Error(), http.StatusInternalServerError)
		}
	}
}

// function validateRequest checks if the method and parameter is correct.
// If the parameter "n" exists and if it is greater than 0.
// If error occur it returns (0, error), else it returns (parameter "n", nil).
func validateRequest(r *http.Request) (int, error) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return 0, fmt.Errorf("%w: %s", ErrMethodIsntAllowed, r.Method)
	}

	key, exist := r.URL.Query()["n"]

	if !exist {
		return 0, fmt.Errorf("%w: %s", ErrParameterIsntProvided, "n")
	}

	length, err := strconv.Atoi(key[0])

	if err != nil {
		return 0, err
	}

	if length < 1 {
		return length, fmt.Errorf("%w: %v", ErrNotPositiveArgument, length)
	}

	return length, nil
}
