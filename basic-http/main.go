package main

import (
	"basic-http/handler"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler.RootHandler)
	mux.HandleFunc("/hello", handler.HelloHandler)
	mux.HandleFunc("/bagus", handler.BagusHandler)
	mux.HandleFunc("/product", handler.ProductHandler)

	log.Println("Starting web on port 8080")

	err := http.ListenAndServe(":8080", mux)
	log.Fatal(err)
}
