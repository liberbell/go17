package main

import (
	"net/http"

	"demo.com/demo"
	"demo.com/demo/html"
)

func main() {
	http.Handle("/person/html", &html.PersonHandler{demo.Person{"Bob", "Smith"}})
	http.Handle("/car/html", &html.CarHandler{demo.Car{"Ford", "Fiesta", 2005}})
}
