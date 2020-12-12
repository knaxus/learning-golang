package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Cases
type Cases struct {
	Statewise []Statewise `json:"statewise"`
}

// Statewise
type Statewise struct {
	ActivePatients    string `json:"active"`
	ConfiremdPatients string `json:"confirmed"`
	Deaths            string `json:"deaths"`
}

func main() {
	//Get the response, otherwise return error
	response, err := http.Get("https://api.covid19india.org/data.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	//Get the response data via reading the response body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	//Unmarshelling the responseData
	var responseObject Cases
	json.Unmarshal(responseData, &responseObject)

	//Results for the covid data
	for i := 0; i < len(responseObject.Statewise); i++ {
		fmt.Println("Active Patients:    ", responseObject.Statewise[i].ActivePatients)
		fmt.Println("Confirmed Patients: ", responseObject.Statewise[i].ConfiremdPatients)
		fmt.Println("Confirmed Deaths:   ", responseObject.Statewise[i].Deaths)
	}
}
