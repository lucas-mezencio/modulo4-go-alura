package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-rest/database"
	"go-rest/models"
	"log"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func GetAllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	err := json.NewEncoder(w).Encode(p)
	if err != nil {
		log.Panicln(err)
	}
}

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var personality models.Personalidade
	database.DB.First(&personality, id)
	json.NewEncoder(w).Encode(personality)
}
