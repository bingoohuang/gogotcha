package main

import (
	"fmt"
)

type noCopy struct{}

func (*noCopy) Lock()   {}
func (*noCopy) Unlock() {}

type Demo struct {
	noCopy noCopy
}

func Copy(d Demo) {
	CopyTwice(d)
}
func CopyTwice(d Demo) {}

func main() {
	d := Demo{}
	fmt.Printf("%+v", d)

	Copy(d)

	fmt.Printf("%+v", d)
}
