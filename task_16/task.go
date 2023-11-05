package main

import "fmt"

func partitionHoare(arr []int, low, high int) int {
	pivot := arr[(high+low)/2]
	i := low
	j := high

	for {
		for arr[i] < pivot {
			i++
		}
		for arr[j] > pivot {
			j--
		}
		if i >= j {
			return j
		}
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		p := partitionHoare(arr, low, high)
		arr = quickSort(arr, low, p)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func main() {
	arr := []int{15, 6, 27, 2, 10, 0, 92}
	fmt.Println(quickSort(arr, 0, len(arr)-1))
}
