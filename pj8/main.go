package main

import (
	"fmt"
	"log"
	"net/http"
)

func CustomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler.")
}

func LogRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("About to serve request.")
		handler.ServeHTTP(w, r)
	})
}

func main() {
	serveMux := http.ServeMux()
	serveMux.HandleFunc("/custom/", CustomHandler)
}
