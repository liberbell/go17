package main

import (
	"io"
	"net/http"
	"net/http/httptest"
)

func main() {
	testHTTPCall := func(responseWriter http.ResponseWriter, request *http.Request) {
		io.WriteString(responseWriter, "This is a test response.")
	}

	request := httptest.NewRequest("GET", "http://testurl.com", nil)
	recorder := httptest.NewRecorder()
	testHTTPCall(recorder, request)
}
