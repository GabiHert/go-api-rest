package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/alura/go-api-rest/interfaces"
	"github.com/alura/go-api-rest/logger"
)

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	logger.Info("controllers.AllPersonalities", "Process started")
	personalities, err := interfaces.IDatabase.GetPersonalities()
	if err != nil {
		json.NewEncoder(w).Encode(err.Error())
	}
	json.NewEncoder(w).Encode(personalities)
	logger.Info("controllers.AllPersonalities", "Process finished")

}
