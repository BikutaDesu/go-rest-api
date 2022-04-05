package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bikutadesu/go-rest-api/database"
	"github.com/bikutadesu/go-rest-api/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	var p []models.Person
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person

	database.DB.First(&person, id)
	json.NewEncoder(w).Encode(person)
}

func InsertPerson(w http.ResponseWriter, r *http.Request) {
	var person models.Person
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Create(&person)
	json.NewEncoder(w).Encode(person)
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person

	database.DB.Delete(&person, id)
	json.NewEncoder(w).Encode(person)
}

func EditPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var person models.Person
	database.DB.First(&person, id)
	json.NewDecoder(r.Body).Decode(&person)
	database.DB.Save(&person)
	json.NewEncoder(w).Encode(person)
}
