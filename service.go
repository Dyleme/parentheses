package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func generateBrackets(length int) string {
	var bracket = map[int]rune{
		0: ')',
		1: ']',
		2: '}',
		3: '(',
		4: '[',
		5: '{',
	}

	rand.Seed(time.Now().Unix())

	var sb strings.Builder

	sb.Grow(length)

	for i := 0; i < length; i++ {
		sb.WriteRune(bracket[rand.Intn(len(bracket))])
	}

	return sb.String()
}

// handler handels only GET method
// Writes sequence of brackets.
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Only method GET is supported")

		return
	}

	q := r.URL.Query()
	key := q.Get("n")

	if key == "" {
		log.Printf("wrong request, parametr doesn't exist")

		return
	}

	length, err := strconv.Atoi(key)

	if err != nil {
		log.Printf("can't convert key %q to int", key)

		return
	}

	fmt.Fprint(w, generateBrackets(length))
}

func main() {
	http.HandleFunc("/generate", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
