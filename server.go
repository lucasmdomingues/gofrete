package main

import (
	"gofrete/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8079"
	}

	r := mux.NewRouter()

	r.HandleFunc("/", handlers.HandlerHome)
	r.HandleFunc("/frete", handlers.FreteHandler).Methods("GET")

	http.ListenAndServe(":"+port, r)
}
