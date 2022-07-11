package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alura/go-api-rest/interfaces"
	"github.com/alura/go-api-rest/logger"
	"github.com/alura/go-api-rest/models"
)

func EditPersonality(w http.ResponseWriter, r *http.Request) {
	logger.Info("controllers.EditPersonality", "Process started")

	vars := mux.Vars(r)
	id := vars["id"]
	var personality models.Personality
	json.NewDecoder(r.Body).Decode(&personality)
	logger.Info("controllers.EditPersonality", "Process started", "id: ", id)

	err := interfaces.IDatabase.EditPersonality(personality, id)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	logger.Info("controllers.EditPersonality", "Process finished")

}
