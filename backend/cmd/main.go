package main

import (
	"backend/internal/adapter/app"
	"backend/internal/adapter/controller/api/server"
	"log"
)

func main() {
	mainApp, err := app.New()
	if err != nil {
		log.Fatalf("failed to create app: %v", err)
	}
	server.Setup(mainApp)
	mainApp.Start()
}
