package main

import (
	"log"
	"os"

	"github.com/RaevNur/simple-go-microservices-api-gateway/api/crud"
	"github.com/RaevNur/simple-go-microservices-api-gateway/api/parser"
	"github.com/RaevNur/simple-go-microservices-api-gateway/configs"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := configs.LoadConfig()
	if err != nil {
		log.Println("Can't load configs:", err)
		os.Exit(1)
	}

	r := gin.Default()

	parser.RegisterRoutes(r, &c)
	crud.RegisterRoutes(r, &c)

	r.Run(c.Port)
}
