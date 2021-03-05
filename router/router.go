package router

import (
	"academy/controller"

	"github.com/gorilla/mux"
)

// InitServer will execute in the defined port
func InitServer() mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/jokes", controller.JokesGetJokes).Methods("GET")
	router.HandleFunc("/api/v1/jokes/{id}", controller.JokesGetOneJoke).Methods("GET")
	router.HandleFunc("/api/v1/update-jokes", controller.UpdateGetData).Methods("GET")
	return *router
}
