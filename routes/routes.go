package routes

import (
	"github.com/gorilla/mux"
	"go-rest/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonality).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
