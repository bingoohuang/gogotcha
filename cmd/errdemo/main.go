package main

import (
	"errors"
	"fmt"
)

func main() {
	еrr := errors.New("foo")
	var err error = nil
	if еrr != nil {
		fmt.Printf("%T %v", err, err)
	}
}
