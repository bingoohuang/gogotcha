package a

import (
	"fmt"
	"importcycle/b"
)

type A struct {
}

func (a A) DoSomethingWithA() {
	fmt.Println(a)
}
func CreateA() *A {
	a := &A{}
	return a
}
func invokeSomethingFromB() {
	o := b.CreateB()
	o.doSomethingWithB()
}
