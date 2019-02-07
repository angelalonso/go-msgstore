package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Message struct {
	Origin  string `json:"origin"`
	Content string `json:"content",omitempty`
}

var msgs []Message

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/msg/{content}", AddMessage).Methods("POST")
	router.HandleFunc("/msg", GetMessage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	var item Message
	item, msgs = msgs[0], msgs[1:]
	json.NewEncoder(w).Encode(item)
	return
}

func AddMessage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var message Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	message.Content = params["content"]
	fmt.Println(message.Content)
	msgs = append(msgs, message)
	json.NewEncoder(w).Encode(msgs)
}
