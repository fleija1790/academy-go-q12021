package router

import (
	"log"
	"net/http"

	"academy/controller"

	"github.com/gorilla/mux"
)

// InitServer will execute in the defined port
func InitServer() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/jokes", controller.JokesGetJokes).Methods("GET")
	router.HandleFunc("/api/v1/jokes/{id}", controller.JokesGetOneJoke).Methods("GET")
	router.HandleFunc("/api/v1/update-jokes", controller.UpdateGetData).Methods("GET")

	log.Fatal(http.ListenAndServe("localhost:3000", router))
}
