package controller

import (
	"academy/model"
	"academy/services/dataload"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//JokesGetJokes all Jokes in the data
func JokesGetJokes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataload.ReadData())
	log.Println("[Info] Returned all Jokes")
}

//JokesGetOneJoke only one joke by ID
func JokesGetOneJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	jokeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.ApiError{
			Error:   "BadRequest",
			Message: "Search are not allow by text or other values, please use an integer",
		})
		log.Println("[Error] not valid ID parameter provided:", err)
		return
	}

	if jokeID > len(dataload.ReadData()) {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(model.ApiError{
			Error:   "NotFound",
			Message: "This joke doesn't exists",
		})
		log.Println("[Error] Id parameter provided is greater than the list of Jokes :", jokeID)
		return
	}

	for _, joke := range dataload.ReadData() {
		if joke.ID == jokeID {
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(joke)
			log.Println("[Info] Returned Joke, ", jokeID)
		}
	}
}
