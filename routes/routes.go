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
	r.HandleFunc("/api/personalidades", controllers.GetAllPersonalities).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.GetPersonality).Methods("Get")
	r.HandleFunc("/api/personalidades", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/api/personalidades/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/api/personalidades/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
