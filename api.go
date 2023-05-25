package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hola Mundo")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	http.ListenAndServe(":8080", mux)
}
