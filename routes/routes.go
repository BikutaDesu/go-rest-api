package routes

import (
	"log"
	"net/http"

	"github.com/bikutadesu/go-rest-api/controllers"
	"github.com/bikutadesu/go-rest-api/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	corsOption := handlers.AllowedOrigins([]string{"*"})
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/people", controllers.InsertPerson).Methods("Post")
	r.HandleFunc("/api/people/{id}", controllers.DeletePerson).Methods("Delete")
	r.HandleFunc("/api/people/{id}", controllers.EditPerson).Methods("Put")
	r.HandleFunc("/api/people", controllers.GetPeople).Methods("Get")
	r.HandleFunc("/api/people/{id}", controllers.GetPerson).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(corsOption)(r)))
}
