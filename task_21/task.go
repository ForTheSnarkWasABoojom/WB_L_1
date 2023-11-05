package main

import (
	"fmt"
)

type Target interface {
	Operation()
}

type Adaptee struct {
}

func (adaptee *Adaptee) AdaptedOperation() {
	fmt.Println("This is method from adaptee!")
}

type Adapter struct {
	*Adaptee
}

func (adapter *Adapter) Operation() {
	adapter.AdaptedOperation()
}

func NewAdapter(adaptee *Adaptee) Target {
	return &Adapter{adaptee}
}

func main() {
	adaptee := Adaptee{}
	adapter := NewAdapter(&adaptee)
	adapter.Operation()
}
