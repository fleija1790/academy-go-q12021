package controller

import (
	"academy/model"
	"academy/services/client"
	"academy/services/updatefile"
	"encoding/json"
	"net/http"
)

//UpdateGetData Updates the CSV from the target API
func UpdateGetData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	jsonData, ret := client.ConsultExternalService()
	var apiError model.ApiError
	if ret != apiError {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(ret)
	} else {
		updatefile.UpdateFile(jsonData)
		w.WriteHeader(http.StatusOK)
		var success model.UpdateOk
		success.Message = "Sucessfully retrieved more Jokes"
		json.NewEncoder(w).Encode(success)
	}

}
