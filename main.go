package main

import (
	"encoding/json"
	"log"
	"net/http" //for server
	"github.com/gorilla/mux" //for handler
)

// Structs for data
type User struct {
	UserID    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Message struct {
	ID      string `json:"id"`
	Message string `json:"message"`
	User    *User  `json:"user"`
}


var messages = make(map[string]Message)

func main() {

	// Appending some data
	messages["1"] = Message{ID: "1", Message: "Hello", User: &User{UserID: "1", FirstName: "Mishaal", LastName: "Naeem", Email: "mishaal@something.com"}}
	messages["2"] = Message{ID: "2", Message: "How are you?", User: &User{UserID: "2", FirstName: "John", LastName: "Doe", Email: "john.doe@something.com"}}

	//router
	r:= mux.NewRouter()

	//handle funcs
	r.HandleFunc("/messages", getMessages).Methods("GET")
	r.HandleFunc("/messages/{id}", newMessage).Methods("POST")
	r.HandleFunc("/messages/{id}", updateMessage).Methods("PUT")
	r.HandleFunc("/messages/{id}", deleteMessage).Methods("DELETE")

	//server
	err:= http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

}

func getMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func newMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)
	var message Message
	_ = json.NewDecoder(r.Body).Decode(&message)
	messages[params["id"]] = message
	json.NewEncoder(w).Encode(message)
}

func updateMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params:= mux.Vars(r)

	if _, ok := messages[params["id"]]; ok {
		var message Message
		_ = json.NewDecoder(r.Body).Decode(&message)
		messages[params["id"]] = message
		json.NewEncoder(w).Encode(message)
	} else {
		log.Panic("id not found")
	}
}

func deleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contet-Type", "application/json")
	params:= mux.Vars(r)
	if msg, ok := messages[params["id"]]; ok {
		delete(messages, params["id"])
		json.NewEncoder(w).Encode(msg)
	} else {
		log.Panic("id not found")
	}
}