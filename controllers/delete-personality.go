package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alura/go-api-rest/interfaces"
	"github.com/alura/go-api-rest/logger"
)

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	logger.Info("controllers.DeletePersonality", "Process started")

	vars := mux.Vars(r)
	id := vars["id"]

	logger.Info("controllers.DeletePersonality", "Process started", "id: ", id)

	err := interfaces.IDatabase.DeletePersonality(id)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	logger.Info("controllers.DeletePersonality", "Process finished")

}
