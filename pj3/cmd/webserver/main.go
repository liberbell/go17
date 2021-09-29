package main

import (
	"net/http"
	"demo.com/demo"
	"demo.com/demo/html"
	"demo.com/demo/plaintext"
)

func main() {
	http.Handle("/person/html", &html.PersonHandler{demo.Person{"Bob, "Smith"}})
	http.Handle("/car/html", )
}