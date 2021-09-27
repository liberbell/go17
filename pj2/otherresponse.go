package main

import (
	"net/http"
)

func main() {
	testServer := httptest(http.HandleFunc(func(resonseWriter http.ResponseWriter, r http.Request) {
		resonseWriter.WriteHeader(http.StatusNotFound)
	}))
}
