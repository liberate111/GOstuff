package main

import (
	"io"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/welcome", welcome)
	mux.HandleFunc("/page", page)

	http.ListenAndServe(":5050", mux)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Welcome!")
}

func page(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "showpage")
}

// localhost:5050/
