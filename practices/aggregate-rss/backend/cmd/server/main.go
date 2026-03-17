package main

import (
	"log"

	"github.com/dungtc/aggregate-rss/backend/internal/api"
	"github.com/dungtc/aggregate-rss/backend/internal/config"
	"github.com/dungtc/aggregate-rss/backend/internal/database"
	"github.com/dungtc/aggregate-rss/backend/internal/scheduler"
)

func main() {
	config.LoadConfig()

	database.InitDB()

	s, err := scheduler.InitCron()
	if err != nil {
		log.Fatal("Failed to setup scheduler: ", err)
	}
	defer s.Shutdown()

	r := api.SetupRouter()

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to run server: ", err)
	}
}
