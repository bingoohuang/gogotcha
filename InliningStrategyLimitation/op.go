package main

func add(a, b float32) float32 {
	if b < 0 {
		panic(`Do not add negative number`)
	}

	return a + b
}

func sub(a, b float32) float32 {
	return a - b
}
