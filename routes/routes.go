package routes

import (
	"log"
	"net/http"

	"github.com/alura/go-api-rest/controllers"
	"github.com/alura/go-api-rest/infra/middleware"
	"github.com/alura/go-api-rest/logger"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	logger.Info("HandleRequest", "Process started")

	r.Use(middleware.Middleware)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/personalities", controllers.CreatePersonality).Methods("Post")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonality).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.DeletePersonality).Methods("Delete")
	r.HandleFunc("/personalities/{id}", controllers.EditPersonality).Methods("Put")
	log.Fatal(http.ListenAndServe(":8080", r))

	logger.Info("HandleRequest", "Process finished")

}
