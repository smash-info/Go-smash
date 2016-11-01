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
	r.HandleFunc("/articles", ArticlesHandler)
	fmt.Println("/ – is available")
	fmt.Println("/products – is available")
	fmt.Println("/articles – is available")
	http.Handle("/", r)
}
