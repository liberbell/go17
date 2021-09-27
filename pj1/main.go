package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, _ := http.Get("https://postman-echo.com/get")

	responseBody, _ := ioutil.ReadAll(response.Body)

	fmt.Println(response.Status)
	fmt.Println(string(responseBody))
}
