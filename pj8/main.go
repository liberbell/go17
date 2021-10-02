package main

import (
	"fmt"
	"net/http"
)

func CustomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler.")
}

func LogRequests(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request))
}

func main() {
	serveMux := http.ServeMux()
	serveMux.HandleFunc("/custom/", CustomHandler)
}
