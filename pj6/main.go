package main

import (
	"fmt"
	"net/http"
)

type TestHandler struct{}

func (h *TestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a custom handler called TestHandler.")
}

func main() {
	htt.Handle
}
