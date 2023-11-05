package main

import "fmt"

type Human struct {
}

func (h *Human) Speak() {
	fmt.Println("Blah blah blah!")
}

type Action struct {
	Human
}

func main() {
	human := Human{}
	action := Action{Human: Human{}}

	human.Speak()
	action.Speak()
}
