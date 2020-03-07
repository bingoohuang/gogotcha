package internstrings

import (
	"fmt"
	"reflect"
	"unsafe"
)

// https://commaok.xyz/post/intern-strings/

// nolint
func pointer(s string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

// nolint
func main() {
	b := []byte("hello")
	s := string(b)
	t := string(b)

	fmt.Println(pointer(s), pointer(t)) // 824634191624 824634191592

	b = []byte("h")
	s = string(b)
	t = string(b)

	fmt.Println(pointer(s), pointer(t)) // 18260008 18260008
}
