package main

import (
	_ "expvar"
	"net/http"
)

func main() {
	http.ListenAndServe(":8100", nil)
}
