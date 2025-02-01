package main

import (
	"log"

	"github.com/oyen-bright/MyTradePad/backend/database"
	"github.com/oyen-bright/MyTradePad/config"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()

	if err != nil {
		panic(err)
	}

	// Setup database connection
	db, err := database.Init(&cfg.Database)

	if err != nil {
		panic(err)
	}

	log.Print(db)

}
