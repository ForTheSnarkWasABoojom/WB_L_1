package main

import (
	"fmt"
	"reflect"
)

func getType(v interface{}) string {
	t := reflect.TypeOf(v)
	return t.String()
}

func main() {
	var x interface{}

	x = 21
	fmt.Println(getType(x))

	x = "!dlrow, olleH"
	fmt.Println(getType(x))

	x = true
	fmt.Println(getType(x))

	x = make(chan int)
	fmt.Println(getType(x))
}
