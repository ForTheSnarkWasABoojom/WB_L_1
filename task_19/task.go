package main

import "fmt"

func reverseString(input string) string {
	runes := []rune(input)

	start, end := 0, len(runes)-1
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}

	return string(runes)
}

func main() {
	input := "An oak is a tree. A rose is a flower. A deer is an animal. A sparrow is a bird. Russia is our fatherland. Death is inevitable."
	fmt.Println("Original:", input)
	reversed := reverseString(input)
	fmt.Println("Reversed:", reversed)
}
