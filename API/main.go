package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Person struct {
	ID      string   `json:"id,omitempty"`
	Fname   string   `json:"fname,omitempty"`
	Lname   string   `json:"lname,omitempty"`
	Address *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", Fname: "Gopikkrishnan", Lname: "Annamalaichamy", Address: &Address{City: "Brighton", State: "Brighton"}})
	people = append(people, Person{ID: "2", Fname: "Keshya", Lname: "Gopikkrishnan", Address: &Address{City: "Brighton", State: "Brighton"}})
	router.HandleFunc("/people", GetPeopleEndpoint).Methods("Get")
	router.HandleFunc("/people/{id}", GetPersonEndpoint).Methods("Get")
	router.HandleFunc("/people/{id}", CreatePersonEndpoint).Methods("Put")
	router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("Delete")
	log.Fatal(http.ListenAndServe(":12345", router))

}
