package main

import "net/http"

func main() {
	fs := http.FileServer(http.Dir("/Users/virtualenv/go17/pj5"))

	http.Handle("/projectfiles/", http.StripPrefix("/projectfiles", fs))

	http.HandleFunc("/samplepdf", func(res http.ResponseWriter, req *http.Request) {
		http.ServeFile(res, req, "/Users/virtualenv/go17/pj5/pdf/matplotlib.pdf")
	})

	http.ListenAndServe(":8001", nil)
}
