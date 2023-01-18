package main

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}
type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Emplyoyee   User   `json:"employee"`
}

func addItem(q http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("Content-Type", "application/json")
	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(profiles)

}
func getAllProfiles(q http.ResponseWriter, r *http.Request) {
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)
}
func getProfile(q http.ResponseWriter, r *http.Request) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to Integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specified ID"))
		return
	}
	profile := profiles[id]
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profile)

}
func updateprofile(q http.ResponseWriter, r *http.Request) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to Integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specified ID"))
		return
	}
	var updateprofile Profile
	json.NewDecoder(r.Body).Decode(&updateprofile)
	profiles[id] = updateprofile
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(updateprofile)

}
func deleteprofile(q http.ResponseWriter, r *http.Request) {
	var idparam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idparam)
	if err != nil {
		q.WriteHeader(400)
		q.Write([]byte("ID could not be converted to Integer"))
		return
	}
	if id >= len(profiles) {
		q.WriteHeader(404)
		q.Write([]byte("no profile found with specified ID"))
		return
	}
	profiles = append(profiles[:id], profiles[:id+1]...)
	q.WriteHeader(200)

}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")
	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")
	router.HandleFunc("/profiles/{id}", getProfile).Methods("GET")
	router.HandleFunc("/profiles/{id}", updateprofile).Methods("PUT")
	router.HandleFunc("/profiles/{id}", deleteprofile).Methods("DELETE")

	http.ListenAndServe(":9000", router) //first pass local host port num then scond thing pass mux router

}
