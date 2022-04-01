package main

// @title Good Food Gateway documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/v1/

import (
	_ "github.com/MSI-GoodFood/GATEWAY/_docs"
	"github.com/MSI-GoodFood/GATEWAY/gateway"

	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil { log.Fatal("Error loading .env file") }

	s := gateway.NewService(os.Getenv("REDIS_URI"), os.Getenv("POSTGRESQL_URI"))

	g := gin.Default()
	s.SetupRoute(g)

	g.Run()
}
