package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/bingoohuang/gogotcha/cgroup"
)

const (
	MB = 1024 * 1024
)

/*
[root@fs03-192-168-126-18 bingoohuang]# ./cgroupdemo
INFO[0000] write tasks with data 5728
INFO[0000] write memory.limit_in_bytes with data 30000000
Child pid is 5728
Alloc = 1 MiB	Sys = 69 MiB
Alloc = 3 MiB	Sys = 69 MiB
Alloc = 7 MiB	Sys = 69 MiB
Alloc = 4 MiB	Sys = 69 MiB
Alloc = 12 MiB	Sys = 69 MiB
Alloc = 12 MiB	Sys = 69 MiB
Alloc = 12 MiB	Sys = 69 MiB
Alloc = 12 MiB	Sys = 69 MiB
已杀死
*/

func main() {
	_ = cgroup.Cgroup{Memory: "30M"}.Limit()

	blocks := make([][MB]byte, 0)

	fmt.Println("Child pid is", os.Getpid())

	for i := 0; ; i++ {
		// nolint:staticcheck
		blocks = append(blocks, [MB]byte{})

		printMemUsage()

		time.Sleep(time.Second)
	}
}

func printMemUsage() {
	var m runtime.MemStats

	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tSys = %v MiB \n", bToMb(m.Sys))
}

func bToMb(b uint64) uint64 {
	return b / MB
}
