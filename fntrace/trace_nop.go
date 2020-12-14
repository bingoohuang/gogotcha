// +build !trace

package fntrace

func noop()         {}
func Trace() func() { return noop }
