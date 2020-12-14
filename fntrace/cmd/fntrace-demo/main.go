package main

import "github.com/bingoohuang/fntrace"

func main() {
	a()
}

func a() {
	defer fntrace.Trace()()
	b()
}

func b() {
	defer fntrace.Trace()()
	c()
}

func c() {
	defer fntrace.Trace()()
	d()
}

func d() {
	defer fntrace.Trace()()
}
