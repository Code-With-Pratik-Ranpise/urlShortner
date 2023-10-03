package main

import (
	"awesomeProject1/pkg/service"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	shortener := service.NewShortener()

	router := mux.NewRouter()
	router.HandleFunc("/shorten", shortener.ShortenURL).Methods("POST")
	router.HandleFunc("/r/", shortener.RedirectURL).Methods("GET")
	router.HandleFunc("/metrics", shortener.Metrics).Methods("GET")

	port := 8080
	fmt.Printf("Server is listening on port %d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	if err != nil {
		return
	}
}
