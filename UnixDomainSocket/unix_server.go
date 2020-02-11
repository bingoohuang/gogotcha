package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"os/signal"
	"syscall"

	"github.com/Masterminds/goutils"
)

func unitHTTPServer(unixPath string) {
	fmt.Println("Unix HTTP server")

	unixListener, err := net.Listen("unix", unixPath)
	if err != nil {
		panic(err)
	}
	defer unixListener.Close()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
			dump, _ := httputil.DumpRequest(req, true)
			abbr, _ := goutils.Abbreviate(string(dump), 900)
			fmt.Printf("%s\n", abbr)
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte("Hello\n"))
		})

		svr := &http.Server{Handler: http.DefaultServeMux}

		fmt.Printf("listened on unix %s\n", unixPath)

		err = svr.Serve(unixListener)
		if err != nil {
			fmt.Println(err)
		}
	}()

	// Use a buffered channel so we don't miss any signals
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}
