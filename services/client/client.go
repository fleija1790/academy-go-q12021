package client

import (
	"academy/model"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//ConsultExternalService will call the Jokes API
func ConsultExternalService() ([]model.Joke, model.ApiError) {
	url := "https://official-joke-api.appspot.com/random_ten"
	var ret model.ApiError
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		ret.Error = "NotFound"
		ret.Message = "Error consulting the external API"
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		// fmt.Println("Error fetching response")
		ret.Error = "Not Found"
		ret.Message = "Error fetching the response"
	}

	defer response.Body.Close()
	data, _ := ioutil.ReadAll(response.Body)

	// Unmarshal JSON data
	var jsonData []model.Joke
	dataErr := json.Unmarshal([]byte(data), &jsonData)
	if dataErr != nil {
		ret.Error = "NotFound"
		ret.Message = "Unable to load json from External request"
	}
	return jsonData, ret
}
