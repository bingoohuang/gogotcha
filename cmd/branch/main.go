package main

import (
	"flag"
	"fmt"
)

func main() {
	tests := make([]bool, 8)
	num := flag.Uint("v", 0, "flags to extract")
	flag.Parse()
	// input value, in binary
	var input uint8 = uint8(*num)
	fmt.Printf("val: %b\n", input)
	// set each boolean to a bit in the input
	if input&(1<<0) != 0 {
		tests[0] = true
	}
	if input&(1<<1) != 0 {
		tests[1] = true
	}
	if input&(1<<2) != 0 {
		tests[2] = true
	}
	if input&(1<<3) != 0 {
		tests[3] = true
	}
	if input&(1<<4) != 0 {
		tests[4] = true
	}
	if input&(1<<5) != 0 {
		tests[5] = true
	}
	if input&(1<<6) != 0 {
		tests[6] = true
	}
	if input&(1<<7) != 0 {
		tests[7] = true
	}
	fmt.Printf("result: %v\n", tests)
	for i, val := range tests {
		fmt.Printf("result %d: %t\n", i, val)
	}
}
