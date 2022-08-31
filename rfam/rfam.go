package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
    "encoding/json"
)

type Result struct {
	Comment		string `json:"comment"`
	Description	string `json:"description"`
}

type Response struct {
	Rfam	Result `json:"rfam"`
}

func main() {
	response, err := http.Get("http://rfam.org/family/RF00360?content-type=application/json")
	if err != nil {
		log.Fatal(err)
	}

	rawData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var r Response 
	json.Unmarshal(rawData, &r)
	fmt.Println("Comment: ", r.Rfam.Comment)
	fmt.Println("Description: ", r.Rfam.Description)
}

