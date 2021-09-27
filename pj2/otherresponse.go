package main

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	testServer := httptest.NewServer(http.HandleFunc(func(resonseWriter http.ResponseWriter, r *http.Request) {
		resonseWriter.WriteHeader(http.StatusNotFound)
		io.WriteString(resonseWriter, "This is a test response.")
	}))
}
