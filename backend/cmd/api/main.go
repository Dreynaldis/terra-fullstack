package main

import (
	"backend/internal/config"
	"backend/internal/database"
	"backend/internal/server"
	"log"
)

func main() {
	cfg := config.LoadConfig()

	dbService := database.NewDatabase(cfg)
	defer func() {
		if err := dbService.Close(); err != nil {
			log.Printf("Error closing database: %v\n", err)
		}
	}()

	srv := server.NewServer(cfg, dbService)
	if err := srv.Run(); err != nil {
		log.Fatalf("Error running server: %v\n", err)
	}
}
