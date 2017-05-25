/*
It should return a map of the counts of each “word” in the string s. The
wc.Test function runs a test suite against the provided function and prints
success or failure.
*/

package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

// WordCount checks for distinct words in a string and returns their count as map
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
