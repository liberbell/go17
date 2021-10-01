package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("/pdf"))

	http.Handle("/projectfiles/", http.StripPrefix("/projectfiles", fs))

	http.HandleFunc("samplepdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "/pdf/matplotlib.pdf")
	})

	http.ListenAndServe(":8001", nil)
}
