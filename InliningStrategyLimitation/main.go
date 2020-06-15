package main

import "fmt"

func main() {
	n := []float32{120.4, -46.7, 32.50, 34.65, -67.45}
	fmt.Printf("The total is %.02f\n", sum(n))
}

func sum(s []float32) float32 {
	var t float32
	for _, v := range s {
		if t < 0 {
			t = add(t, v)
		} else {
			t = sub(t, v)
		}
	}

	return t
}

// nolint:unused, gochecknoglobals
var a int

// Aadd demos SSA dump
// nolint:unused,deadcode
func Aadd() {
	a++
}
