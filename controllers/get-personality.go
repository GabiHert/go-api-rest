package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/alura/go-api-rest/interfaces"
	"github.com/alura/go-api-rest/logger"
)

func GetPersonality(w http.ResponseWriter, r *http.Request) {
	logger.Info("controllers.GetPersonality", "Process started")
	vars := mux.Vars(r)
	id := vars["id"]

	logger.Info("controllers.GetPersonality", "Process started", "id: ", id)

	personality, err := interfaces.IDatabase.GetPersonality(id)

	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	personalityJson, _ := json.Marshal(personality)

	logger.Info("controllers.GetPersonality", "Process finished", "personality:", string(personalityJson))

	json.NewEncoder(w).Encode(personality)

}
