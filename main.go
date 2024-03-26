package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var profiles []Profile = []Profile{}

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

type Profile struct {
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee    User   `json:"employee"`
}

func addItem(w http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	w.Header().Set("Content-Type", "application/json")
	profiles = append(profiles, newProfile)

	json.NewEncoder(w).Encode(profiles)
}

func getAllProfiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profiles)
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer."))
		return
	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("No profile found with specified ID"))
		return
	}
	profile := profiles[id]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

func updateProfile(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("No profile found with specified ID"))
		return
	}

	var updateProfile Profile
	json.NewDecoder(r.Body).Decode(&updateProfile)

	profiles[id] = updateProfile

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updateProfile)
}

func deleteProfile(w http.ResponseWriter, r *http.Request) {
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to Integer"))
		return
	}

	if id >= len(profiles) {
		w.WriteHeader(404)
		w.Write([]byte("No profile found with specified ID"))
		return
	}
	profiles = append(profiles[:id], profiles[:id+1]...)

	w.WriteHeader(200)
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")

	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")

	router.HandleFunc("/profiles/{id}", getProfile).Methods("GET")

	router.HandleFunc("/profiles/{id}", updateProfile).Methods("PUT")

	router.HandleFunc("/profiles/{id}", deleteProfile).Methods("DELETE")
	fmt.Printf("Starting server at port 9000\n")
	http.ListenAndServe(":9000", router)
	//fmt.Println("Listening on port: 9000 ")
}
