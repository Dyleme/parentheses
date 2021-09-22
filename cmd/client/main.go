package main

import (
	"fmt"
	"io"
	"net/http"
	"parentheses/parentheses"
	"strconv"
)

const urlPath = "http://localhost:8080/generate?n="

const requestAmount = 1000

// DoRequest do request to the url with parameter n equals to amount.
// And writes the result of IsBalanced function, which argument is response body, to the channel.
func DoRequest(amount int, results chan bool) {
	if amount < 0 {
		fmt.Println(fmt.Errorf("amount of brackets is not positive %v", amount))
		results <- false
	}

	resp, err := http.Get(urlPath + strconv.Itoa(amount))
	if err != nil {
		fmt.Println(err)
		results <- false
	}
	defer resp.Body.Close()

	body := make([]byte, amount)
	_, err = resp.Body.Read(body)

	if err != io.EOF {
		fmt.Println(err)
		results <- false
	}

	results <- parentheses.IsBalanced(string(body))
}

// BalancedBracketProbability calculate amount of
// balanced bracket sequences which consist of bracketAmount characters.
func BalancedBracketProbability(bracketAmount int) float32 {
	counter := 0
	results := make(chan bool)

	for i := 0; i < requestAmount; i++ {
		go DoRequest(bracketAmount, results)
	}

	for i := 0; i < requestAmount; i++ {
		if <-results {
			counter++
		}
	}

	return float32(counter) / requestAmount
}

func main() {
	for i := 2; i <= 8; i *= 2 {
		fmt.Println(BalancedBracketProbability(i))
	}
}
