package main

import (
	"fmt"
	"net/http"
)

type PersonHandler struct {
	FirstName string
	LastName  string
}

type CarHandler struct {
	Make  string
	Model string
	Year  int
}

func (h *PersonHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Name: %v %v", h.FirstName, h.LastName)
}

func (c *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Make: %v  -- Model: %v -- Year: %v", c.Make, c.Model, c.Year)
}

func main() {
	http.Handle("/person")
}
