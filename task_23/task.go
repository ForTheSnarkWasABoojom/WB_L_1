package main

import "fmt"

func removeElement(slice []int, index int) []int {
	if index < 0 || index >= len(slice) {
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice: ", slice)

	indexToRemove := 2
	slice = removeElement(slice, indexToRemove)

	fmt.Println("Slice after removal: ", slice)
}
