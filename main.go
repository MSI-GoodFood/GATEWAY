package main

// @title Good Food Gateway documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/v1/

import (
	_ "github.com/MSI-GoodFood/GATEWAY/docs"
	"github.com/MSI-GoodFood/GATEWAY/gateway"
	"google.golang.org/grpc"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	API_GESTION_URL := os.Getenv("API_GESTION_URL")
	REDIS_URL := os.Getenv("REDIS_URL")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// Setup GRPC GESTION API
	serverGestion, err := grpc.Dial(API_GESTION_URL, grpc.WithInsecure())
	if err != nil { log.Fatalf("did not connect to API GESTION: %v", err) }
	defer serverGestion.Close()

	s := gateway.NewService(REDIS_URL, DATABASE_URL)

	g := gin.Default()
	s.SetupRoute(g, serverGestion)

	g.Run()
}
