package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func root(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, dear Guest!")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", root)

	http.ListenAndServe(":8000", r)
}
