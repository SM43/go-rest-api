package main

import (
	"log"
	"net/http"
	"github.com/sm43/go-rest-api/lib"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", lib.HomeLink)
	router.HandleFunc("/event", lib.CreateEvent).Methods("POST")
	router.HandleFunc("/events", lib.GetAllEvents).Methods("GET")
	router.HandleFunc("/events/{id}", lib.GetOneEvent).Methods("GET")
	router.HandleFunc("/events/{id}", lib.UpdateEvent).Methods("PATCH")
	router.HandleFunc("/events/{id}", lib.DeleteEvent).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
