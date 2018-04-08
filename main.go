package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type User struct {
	ID        string   `json:"id,omitempty"`
	DeviceId string   `json:"firstname,omitempty"`
	Phone  string   `json:"lastname,omitempty"`
}

// our main function
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", createUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(user)
}