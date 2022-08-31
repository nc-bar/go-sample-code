package main

import (
	"net/http"
	"encoding/json"
	"log"
)

type Message struct {
	Content	string `json:"content"`
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	m := Message{Content: "This is a message!"}
	json.NewEncoder(w).Encode(m)
}

func main() {
	http.HandleFunc("/", showMessage) 
	log.Fatal(http.ListenAndServe(":8000", nil))
}
