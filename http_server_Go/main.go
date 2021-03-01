package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "<h1> Hello %s, I am your server </h1>", r.URL.Path[1:])

}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":7043", nil)
}