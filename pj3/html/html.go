package html

import (
	"fmt"
	"net/http"

	"demo.com/demo"
)

type CarHandler struct {
	Car demo.Car
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<font color='green><b>Make:</b></font>   <i>%v</i><BR><font color='green'><b>Model:</b></font>   <i>%v</i><BR><font color='green'><b>Year:</b></font> <i>%v</i>")
}

type PersonHandler struct {
	Person demo.Person
}

func (h *PersonHandler) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<font color='green'><b>Name:</b></font>   <i>%v</i>", h.Person.FirstName, h.Person.LastName)
}
