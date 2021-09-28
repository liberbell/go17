package html

import (
	"fmt"
	"net/http"
)

type CarHandler struct {
	Car demo.Car
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<font color='green><b>Make:</b></font>   <i>%v</i><BR><font color='green'><b>Model:</b></font>   <i>%v</i>")
}
