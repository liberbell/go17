package planetext

import (
	"net/http"
	"fmt"
	"demo.com/demo"
)

type CarHandler struct {
	car.demo.Car
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r http.Request) {
	fmt.Fprintf(w "Make: %v -- Model: %v -- Year: %v", h.Car.Make, h.Car.Model, h.Car.Year)
}