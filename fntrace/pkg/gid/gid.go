package gid

import (
	"bytes"
	"runtime"
	"strconv"
)

func GetGID() uint64 {
	n, _ := strconv.ParseUint(GetGIDString(), 10, 64)
	return n
}

func GetGIDString() string {
	return string(GetGIDBytes())
}
func GetGIDBytes() []byte {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	return b[:bytes.IndexByte(b, ' ')]
}
