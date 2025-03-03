package main

import (
	"github.com/ardwiinoo/micro-music/users/config"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures"
	"github.com/ardwiinoo/micro-music/users/internal/infrastructures/http"
	"log"
	"strconv"
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
