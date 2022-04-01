package main

// @title Good Food Gateway documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/v1/

import (
	_ "github.com/MSI-GoodFood/GATEWAY/_docs"
	"github.com/MSI-GoodFood/GATEWAY/gateway"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	s := gateway.NewService(os.Getenv("REDIS_URL"), os.Getenv("DATABASE_URL"))

	g := gin.Default()
	s.SetupRoute(g)

	g.Run()
}
