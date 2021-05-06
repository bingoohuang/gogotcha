package main

import (
	"io/ioutil"
	"log"
	"os"
	"runtime"
	"syscall"
)

type File struct{ d int }

func main() {
	file, err := ioutil.TempFile("", "keepalive")
	if err != nil {
		log.Fatal(err)
	}
	file.Write([]byte("keepalive"))
	file.Close()
	defer os.Remove(file.Name())

	p := openFile(file.Name())
	content := readFile(p.d)

	// Ensure p is not finalized until Read returns
	runtime.KeepAlive(p)

	println("Here is the content: " + content)
}

func openFile(path string) *File {
	d, err := syscall.Open(path, syscall.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}

	p := &File{d}
	runtime.SetFinalizer(p, func(p *File) {
		syscall.Close(p.d)
	})

	return p
}

func readFile(descriptor int) string {
	doSomeAllocation()

	var buf [1000]byte
	_, err := syscall.Read(descriptor, buf[:])
	if err != nil {
		panic(err)
	}

	return string(buf[:])
}

func doSomeAllocation() {
	var a *int

	// memory increase to force the GC
	for i := 0; i < 10000000; i++ {
		i := 1
		a = &i
	}

	_ = a
}
