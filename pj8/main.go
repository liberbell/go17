package main

import "net/http"

func CustomHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	serveMux := http.ServeMux()
	serveMux.HandleFunc("/custom/", CustomHandler)
}
