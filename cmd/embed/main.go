package main

import (
	"embed"
	"net/http"
)

//go:embed assets/*
var assets embed.FS

//go:embed index.html
var html []byte

func main() {
	http.Handle("/assets/", http.FileServer(http.FS(assets)))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "text/html")
		w.Write(html)
	})
	http.ListenAndServe(":8080", nil)
}
