package main

import (
	"fmt"
	"net/http"
)

func main() {
	displayStartup()
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products", ProductsHandler)
	fmt.Println("/ – is available")
	fmt.Println("/products – is available")
	http.Handle("/", r)
}
