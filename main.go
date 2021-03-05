package main

import (
	"academy/router"
	"academy/services/dataload"
	"fmt"
	"net/http"
	"os"

	"log"
)

func main() {
	chkLoad := dataload.LoadData()
	if chkLoad != "" {
		fmt.Println(chkLoad)
		os.Exit(1)
	}
	runServer := router.InitServer()

	log.Fatal(http.ListenAndServe("localhost:3000", &runServer))
}
