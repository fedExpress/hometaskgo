package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test-api/api"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/event", api.CreateEvent).Methods("POST")
	router.HandleFunc("/events/{id}", api.GetEvent).Methods("GET")
	router.HandleFunc("/events/{id}", api.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", api.DeleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
