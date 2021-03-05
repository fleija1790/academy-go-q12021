package controller

import (
	"academy/model"
	"academy/services/dataload"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//JokesGetJokes all Jokes in the data
func JokesGetJokes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataload.ReadData())
}

//JokesGetOneJoke only one joke by ID
func JokesGetOneJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	vars := mux.Vars(r)
	jokeID, err := strconv.Atoi(vars["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		var ret model.ApiError
		ret.Error = "BadRequest"
		ret.Message = "Please provide a valid ID"
		json.NewEncoder(w).Encode(ret)
	}

	if len(dataload.ReadData()) <= jokeID {
		w.WriteHeader(http.StatusNotFound)
		var ret model.ApiError
		ret.Error = "NotFound"
		ret.Message = "This joke doesn't exists"
		json.NewEncoder(w).Encode(ret)
	}

	for _, joke := range dataload.ReadData() {
		if joke.ID == jokeID {
			w.WriteHeader(http.StatusFound)
			json.NewEncoder(w).Encode(joke)
		}
	}
}
