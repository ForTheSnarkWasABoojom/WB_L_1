package main

import "fmt"

func main() {
	a, b := true, false
	fmt.Println(a, b)
	b, a = a, b
	fmt.Println(a, b)
}
