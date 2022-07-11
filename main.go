package main

import (
	"github.com/alura/go-api-rest/interfaces"
	"github.com/alura/go-api-rest/logger"
	"github.com/alura/go-api-rest/routes"
)

func main() {
	logger.Info("main", "Process started")

	interfaces.IDatabase.ConnectDB()

	routes.HandleRequest()

	logger.Info("main", "Process finished")
}
