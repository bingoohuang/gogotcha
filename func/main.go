package main

import (
	"fmt"
	"net/http"
)

// Handler handles ServeHTTP.
type Handler interface {
	ServeHTTP(http.ResponseWriter, *http.Request)
}

// HandlerFunc ...
type HandlerFunc func(http.ResponseWriter, *http.Request)

func (h HandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

func main() {
	var h Handler

	h = HandlerFunc(nil)

	fmt.Println(h)
}
