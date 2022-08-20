package main

import (
	"github.com/argalon/routes"
	"go.uber.org/zap"
)

func main() {

	logger := zap.L()
	logger.Info("Argalon Service for divide is starting...")

	logger.Info("Argalon Divide Server Listening for API Calls.")
	r := routes.SetupRouter()

	if err := r.Run(":8088"); err != nil {
		logger.Error("Failed to run server", zap.Error(err))
	}
}
