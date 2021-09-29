package main

import "net/http"

func main() {
	http.Handle("/person/html", &html.PersonHandler{demo.Person{"Bob, "Smith"}})
}