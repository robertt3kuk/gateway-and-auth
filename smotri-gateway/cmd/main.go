package main

import (
	"log"

	"smotri-gateway/pkg/auth"
	"smotri-gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	auth.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
