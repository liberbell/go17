package main

import (
	"fmt"
	"net/http"
)

func NameHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Name: %s %s", p.ByName("firstname"), p.ByName("lastname"))
}

func main() {
	router := httprouter.New()
}
