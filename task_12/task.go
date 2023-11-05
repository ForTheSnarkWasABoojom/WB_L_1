package main

import (
	"fmt"
)

func main() {
	set := make(map[string]bool)

	strings := []string{"cat", "cat", "dog", "cat", "tree"}

	for _, str := range strings {
		set[str] = true
	}

	fmt.Println("Set:")
	for str := range set {
		fmt.Println(str)
	}
}
