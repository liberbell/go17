package html

import "net/http"

type CarHandler struct {
	Car demo.Car
}

func (h *CarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
}
