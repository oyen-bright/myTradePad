package main

import (
	"log"

	"github.com/oyen-bright/MyTradePad/config"
)

// var (
// 	db *gorm.DB
// )

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()

	log.Print(cfg)

	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

}
