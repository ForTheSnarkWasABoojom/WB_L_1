package main

import (
	"fmt"
)

func intersect(arr1, arr2 []int) []int {

	map1 := make(map[int]bool)

	for _, num := range arr1 {
		map1[num] = true
	}

	intersect := []int{}

	for _, num := range arr2 {
		if map1[num] {
			intersect = append(intersect, num)
		}
	}

	return intersect
}

func showArr(arr1 []int, str string) {
	fmt.Println(str)
	for i := range arr1 {
		fmt.Printf("%d ", arr1[i])
	}
}

func main() {
	arr1 := []int{10, 20, 22, 79, 88}
	arr2 := []int{6, 14, 20, 79, 89}

	intersection := intersect(arr1, arr2)

	showArr(arr1, "First array:")
	showArr(arr2, "\nSecond array:")
	fmt.Println("\nIntersection of arrays:", intersection)
}
