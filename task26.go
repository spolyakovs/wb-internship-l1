package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "abcd"
	s2 := "aabcd"
	s3 := "abCdefAf"

	fmt.Printf("String: \"%v\", areSimbolsUnique: %v\n", s1, areSimbolsUnique(s1))
	fmt.Printf("String: \"%v\", areSimbolsUnique: %v\n", s2, areSimbolsUnique(s2))
	fmt.Printf("String: \"%v\", areSimbolsUnique: %v\n", s3, areSimbolsUnique(s3))
}

func areSimbolsUnique(s string) bool {
	runes := []rune(strings.ToLower(s))
	alreadySeenRunes := make(map[rune]bool)

	for _, r := range runes {
		alreadySeen, ok := alreadySeenRunes[r]
		if ok && alreadySeen {
			return false
		}

		alreadySeenRunes[r] = true
	}

	return true
}
