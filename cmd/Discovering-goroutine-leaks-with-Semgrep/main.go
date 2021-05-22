package main

import (
	"fmt"
	"runtime"
	"time"
)

// https://www.trailofbits.com/post/discovering-goroutine-leaks-with-semgrep

func main() {
	requestData(1 * time.Millisecond)
	time.Sleep(time.Second * 1)
	fmt.Printf("Number of hanging goroutines: %d", runtime.NumGoroutine()-1)
}

func requestData(timeout time.Duration) string {
	dataChan := make(chan string) // problem version
	// dataChan := make(chan string, 1) // fixed version

	go func() {
		newData := requestFromSlowServer()
		dataChan <- newData // block, this goroutine may be leaked
	}()
	select {
	case result := <-dataChan:
		fmt.Printf("[+] request returned: %s", result)
		return result
	case <-time.After(timeout):
		fmt.Println("[!] request timeout!")
		return ""
	}
}

func requestFromSlowServer() string {
	time.Sleep(500 * time.Millisecond)
	return "very important data"
}
