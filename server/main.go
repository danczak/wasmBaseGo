package main

import (
	"fmt"
	"net/http"
)

const (
	PORT = 8080
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./wwwroot")))
	http.ListenAndServe(fmt.Sprintf("localhost:%d", PORT), nil)
}

