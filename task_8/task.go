package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func bitTo1(n int, pos int) int {
	n |= (1 << pos)
	return n
}

func bitTo0(n int, pos int) int {
	n &^= (1 << pos)
	return n
}

func main() {
	rand.Seed(time.Now().UnixNano())
	variable := rand.Intn(math.MaxInt64)
	fmt.Printf("Current variable is %d\n", variable)

	fmt.Println("Enter i - bit to be changed and v - the value to which you want to change the bit:")
	var i, v int
	_, err := fmt.Scanln(&i, &v)
	CheckError(err)

	if v == 0 {
		fmt.Println(bitTo0(variable, i))
	} else {
		fmt.Println(bitTo1(variable, i))
	}

}
