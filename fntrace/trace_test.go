// go test -tags trace
package fntrace_test

import "github.com/bingoohuang/fntrace"

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

func ExampleTrace() {
	a()
	// Output:
	// g[01]:->github.com/bingoohuang/fntrace_test.a
	// g[01]:-->github.com/bingoohuang/fntrace_test.b
	// g[01]:--->github.com/bingoohuang/fntrace_test.c
	// g[01]:---->github.com/bingoohuang/fntrace_test.d
	// g[01]:<----github.com/bingoohuang/fntrace_test.d
	// g[01]:<---github.com/bingoohuang/fntrace_test.c
	// g[01]:<--github.com/bingoohuang/fntrace_test.b
	// g[01]:<-github.com/bingoohuang/fntrace_test.a
}
