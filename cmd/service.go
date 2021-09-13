package main

import (
	"log"
	"net/http"

	"parentheses/handlers"
)

func main() {
	http.Handle("/generate", &handlers.GeneratorHandler{handlers.DefaultGenerate})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
