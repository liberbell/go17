package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("/pdf"))

	http.Handle("/projectfiles/", http.StripPrefix("/projectfiles", fs))
}
