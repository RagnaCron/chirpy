package main

import (
	"log"
	"net/http"
)

func main() {
	var server http.Server

	mux := http.NewServeMux()

	server.Addr = ":8080"
	server.Handler = mux

	defer server.Close()

	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
