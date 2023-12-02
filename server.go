package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Print("starting app")
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":80", nil)
}

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello world</h1>"))
}
