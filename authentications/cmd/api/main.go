package main

import (
	"log"
	"strconv"

	"github.com/ardwiinoo/micro-music/authentications/config"
	_ "github.com/ardwiinoo/micro-music/authentications/docs"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/authentications/internal/infrastructures/http"
)

func main() {
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	container, err := infrastructures.NewContainer()

	if err != nil {
		log.Fatalf("Failed to initialize container: %v", err)
	}

	defer container.Close()

	server := http.CreateServer(container)

	port := ":" + strconv.Itoa(config.Cfg.App.Port)
	log.Printf("Starting server on %s", port)
	
	if err := server.Listen(port); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}