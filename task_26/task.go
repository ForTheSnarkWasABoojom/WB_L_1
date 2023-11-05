package main

import (
	"fmt"
)

func toLower(char rune) rune {
	if 'A' <= char && char <= 'Z' {
		return char + ('a' - 'A')
	}
	return char
}

func isUnique(s string) bool {
	uniqueCharacters := make(map[rune]bool)

	for _, char := range s {
		char = toLower(char)
		if uniqueCharacters[char] {
			return false
		}
		uniqueCharacters[char] = true
	}
	return true
}

func main() {
	strA := "Hello, World!"
	strB := "HelLo, WOrd!"
	strC := "Hel, Word!"

	fmt.Printf("String A: %v\n", strA)
	fmt.Printf("String B: %v\n", strB)
	fmt.Printf("String C: %v\n", strC)

	fmt.Printf("Is all characters in string A unique? - %v\n", isUnique(strA))
	fmt.Printf("Is all characters in string B unique? - %v\n", isUnique(strB))
	fmt.Printf("Is all characters in string C unique? - %v\n", isUnique(strC))
}
