package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "snow dog sun"

	fmt.Println("Original:", s)
	fmt.Println("Reversed words:", reverseWordsString(s))
}

// very similar to task 19
func reverseWordsString(s string) string {
	words := strings.Split(s, " ")

	for i := 0; i < len(words)/2; i++ {
		words[i], words[(len(words)-1)-i] = words[(len(words)-1)-i], words[i]
	}

	return strings.Join(words, " ")
}
