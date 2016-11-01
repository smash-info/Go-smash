package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Printf("Starting up server")
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	r.HandleFunc("/articles", ArticlesHandler)
	http.Handle("/", r)
}
