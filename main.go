package main

import (
	embed "embed"
	"log"
	"net/http"
)

//go:embed index.html
var indexPage []byte

//go:embed assets/*
var assets embed.FS

func main() {
	assetsFs := http.FileServer(http.FS(assets))

	mux := http.NewServeMux()
	mux.Handle("/assets/", assetsFs)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write(indexPage)
	})

	log.Println("http://localhost:8000")
	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}
}
