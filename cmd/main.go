package main

import (
	"log"
	"net/http"

	"parentheses/handlers"
	"parentheses/parentheses"
)

// BracketsGenerator is implementation of Generator which generate random sequence of brackets.
type BracketsGenerator struct{}

// Generate function generate random sequence of brackets.
func (g *BracketsGenerator) Generate(length int) string {
	return parentheses.GenerateBrackets(length)
}

func main() {
	genSv := BracketsGenerator{}
	http.Handle("/generate", handlers.New(&genSv))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
