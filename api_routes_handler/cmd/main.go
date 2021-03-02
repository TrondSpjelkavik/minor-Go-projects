package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("This is a text")
}

func main()  {
	fmt.Println("Starting Server")

	router := chi.NewRouter()

	router.Get("/trond/users", getHandler)

	log.Fatal(http.ListenAndServe(":8080", router))

	fmt.Println("Server is Listening")


}