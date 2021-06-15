package main

import (
	"fmt"
	"os"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func IsTTY() bool {
	fi, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return fi.Mode()&os.ModeCharDevice != 0
}

func main() {
	fmt.Println("IsTTY:", IsTTY())

	bar := pb.StartNew(100)
	for i := 0; i < 100; i++ {
		time.Sleep(100 * time.Millisecond)
		bar.Increment()
	}
	bar.Finish()
}
