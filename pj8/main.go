package main

import (
	"fmt"
	"net/http"
)

func CustomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler.")
}

func main() {
	serveMux := http.ServeMux()
	serveMux.HandleFunc("/custom/", CustomHandler)
}
