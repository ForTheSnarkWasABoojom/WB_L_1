package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(4 << 20)
	b := big.NewInt(2 << 20)
	c := new(big.Int)

	fmt.Printf("A = %s, B = %s\n", a.String(), b.String())

	c.Add(a, b)
	fmt.Printf("A + B = %s\n", c.String())

	c.Sub(a, b)
	fmt.Printf("A - B = %s\n", c.String())

	c.Mul(a, b)
	fmt.Printf("A * B = %s\n", c.String())

	c.Div(a, b)
	fmt.Printf("A / B = %s\n", c.String())
}
