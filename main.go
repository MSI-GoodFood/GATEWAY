package main

// @title Good Food Gateway documentation
// @version 1.0.0
// @host localhost:8000
// @BasePath /v1/

import (
	"fmt"
	"gateway/gateway"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	_ "gateway/docs"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := gateway.NewService(os.Getenv("REDIS_URI"), os.Getenv("POSTGRESQL_URI"))

	g := gin.Default()
	s.SetupRoute(g)

	port := os.Getenv("PORT")
	g.Run(fmt.Sprint(":", port))
}
