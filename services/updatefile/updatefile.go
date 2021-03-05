package updatefile

import (
	"academy/model"
	"academy/services/dataload"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readCsvFile(filePath string) ([][]string, []string) {
	var errorMessage []string
	f, err := os.Open(filePath)
	if err != nil {
		errorMessage = append(errorMessage, "Unable to read input file: "+filePath, err.Error())
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		errorMessage = append(errorMessage, "Unable to parse file as CSV for "+filePath, err.Error())
	}

	return records, errorMessage
}

//UpdateFile the CSV from the target API
func UpdateFile(jsonData []model.Joke) {
	pwd, err := os.Getwd()
	csvFile, err := os.OpenFile(pwd+"/data/data.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()
	records, chkRecords := readCsvFile(pwd + "/data/data.csv")
	if len(chkRecords) != 0 {
		fmt.Println("Error Updating file ")
	}

	counter := len(records) + 1
	writer := csv.NewWriter(csvFile)

	for i, usance := range jsonData {
		var row []string
		row = append(row, strconv.Itoa(counter+i))
		row = append(row, strings.Replace(usance.Setup, "\n", "", -1))
		row = append(row, usance.Punchline)
		writer.Write(row)
	}

	writer.Flush()
	dataload.LoadData()
}
