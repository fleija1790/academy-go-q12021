package controller

import (
	"academy/services/client"
	"academy/services/updatefile"
	"net/http"
)

//UpdateGetData Updates the CSV from the target API
func UpdateGetData(w http.ResponseWriter, r *http.Request) {
	jsonData := client.ConsultExternalService()
	updatefile.UpdateFile(jsonData)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
