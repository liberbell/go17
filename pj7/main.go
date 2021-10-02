package main

import "net/http"

type PersonHandler struct {
	FirstName string
	LastName  string
}

type CarHandler struct {
	Make  string
	Model string
	Year  int
}

func main() {
	http.Handle("/person")
}
