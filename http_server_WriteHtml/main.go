package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)

	if err != nil {
		return "", err
	}

	return string(bytes), nil

}

func handler(w http.ResponseWriter, r *http.Request) {

	var html, err = loadFile("index.html")

	if err != nil {
		w.WriteHeader(404)
		fmt.Fprint(w, "404: sorry this page does not exist")
	}

	fmt.Fprint(w, html)

}

func main() {

	http.HandleFunc("/", handler)

	http.ListenAndServe(":7043", nil)
}