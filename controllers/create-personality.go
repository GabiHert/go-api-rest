package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alura/go-api-rest/interfaces"
	"github.com/alura/go-api-rest/logger"
	"github.com/alura/go-api-rest/models"
)

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	logger.Info("controllers.CreatePersonality", "Process started")

	var personality models.Personality

	json.NewDecoder(r.Body).Decode(&personality)

	interfaces.IDatabase.CreatePersonality(personality)

	logger.Info("controllers.CreatePersonality", "Process finished")
}
