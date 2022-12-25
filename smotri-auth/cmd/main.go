package main

import (
	"log"

	"smotri-auth/config"
	"smotri-auth/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %s", err)
	}
	app.Run(cfg)
}
