package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./static/index.html")
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}