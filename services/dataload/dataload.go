package dataload

import (
	"academy/model"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

var allJokes []model.Joke

//LoadData this will load the data from the csv
func LoadData() string {
	var message string
	allJokes = append(allJokes[:0])
	pwd, _ := os.Getwd()
	csvFile, err := os.Open(pwd + "/data/data.csv")
	if err != nil {
		message = "Error openning file: " + err.Error()
	}
	defer csvFile.Close()
	r := csv.NewReader(csvFile)
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err)
		message = "Error reading file: " + err.Error()
		return message
	}
	if len(records) == 0 {
		message = "Currently there are not Jokes Stored... this is not funny!"
		return message
	}
	for _, rec := range records {
		var insert model.Joke
		insert.ID, err = strconv.Atoi(rec[0])
		insert.Setup = rec[1]
		insert.Punchline = rec[2]
		allJokes = append(allJokes, insert)
	}
	return message
}

//ReadData will return the data from the file
func ReadData() []model.Joke {
	return allJokes
}
