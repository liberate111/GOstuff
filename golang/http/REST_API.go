package main

import (
	"encoding/json"
	"net/http"
)

type Book struct {
	Title  string `json:"Title"`
	Author string `json:"Author"`
	Text   string `json:"Text"`
}

func BooksHandler(w http.ResponseWriter, r *http.Request) {
	books := []Book{
		Book{"Book1", "fufuu1", "Chaper1"},
		Book{"Book2", "fufuu2", "Chaper2"},
		Book{"Book3", "fufuu3", "Chaper3"},
	}

	json.NewEncoder(w).Encode(books)

}

func main() {

	http.HandleFunc("/books", BooksHandler)
	http.ListenAndServe(":5051", nil)
}

// localhost:5051/
