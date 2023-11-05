package main

import (
	"fmt"
	"strings"
)

func extractWords(sentence string) []string {
	words := strings.Fields(sentence)
	return words
}

func reverseArr(arr []string) []string {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
func main() {
	sentence := "Ad cogitandum et agendum homo natus est"
	words := extractWords(sentence)
	reverseArr := reverseArr(words)
	reverse := strings.Join(reverseArr, " ")
	fmt.Println(sentence)
	fmt.Println(reverse)
}
