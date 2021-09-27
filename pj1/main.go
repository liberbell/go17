package main

import "net/http"

func main() {
	response, _ := http.Get("https://postman-echo.com/get")
}
