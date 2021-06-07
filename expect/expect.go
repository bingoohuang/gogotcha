// Package expect contains various test assertion helpers.
//
// The functions this package recover the *testing.T value passed
// to the Test from the call stack. This implies that expect
// functions should not be called from a goroutine launched from
// a test. However, this restriction applies to calling t.Fatal/FailNow
// directly, so here we are.
package expect

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"runtime"
	"testing"
	"unsafe"
)

// Nil fails the test if v is not nil.
func Nil(v interface{}) {
	if v == nil {
		return
	}

	if t := getT(); t != nil {
		t.Helper()
		t.Fatalf("expected: %v, got: %v", nil, v)
	} else {
		log.Panicf("%v is not nil", v)
	}
}

// True fails the test if v is not true.
func True(v bool) {
	if v {
		return
	}

	if t := getT(); t != nil {
		t.Helper()
		t.Fatalf("expected: %v, got: %v", true, false)
	} else {
		log.Panicf("%t is not true", v)
	}
}

// getT returns the address of the testing.T passed to testing.tRunner
// which called the function which called getT. If testing.tRunner cannot
// be located in the stack, say if getT is not called from the main test
// goroutine, getT returns nil.
func getT() *testing.T {
	var p uintptr
	var buf [8192]byte
	n := runtime.Stack(buf[:], false)

	const format = "testing.tRunner(%v"
	for sc := bufio.NewScanner(bytes.NewReader(buf[:n])); sc.Scan(); {
		if n, _ := fmt.Sscanf(sc.Text(), format, &p); n == 1 {
			return (*testing.T)(unsafe.Pointer(p))
		}
	}

	return nil
}
