package main

import (
	"io"
	"net/http"
)

func main() {
	testHTTPCall := func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "This is a test response.")
	}
}
