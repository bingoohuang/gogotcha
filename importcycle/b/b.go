package b

import (
	"fmt"
	"importcycle/a"
)

type B struct {
}

func (b B) doSomethingWithB() {
	fmt.Println(b)
}
func CreateB() *B {
	b := B{}
	return &b
}
func invokeSomethingFromA() {
	o := a.CreateA()
	o.DoSomethingWithA()
}
