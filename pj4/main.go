package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/unrecognized_path.html")
	})

	http.HandleFunc("/demopath1/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "html/demopath1.html")
	})
}
