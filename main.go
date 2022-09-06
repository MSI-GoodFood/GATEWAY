package main

// @title Good Food Gateway documentation
// @version 1.0.0
// @host localhost:8080
// @BasePath /api/v1/

import (
	_ "github.com/MSI-GoodFood/GATEWAY/docs"
	"github.com/MSI-GoodFood/GATEWAY/gateway"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	_ = godotenv.Load()

	API_GESTION_URL := os.Getenv("API_GESTION_URL")
	API_ORDER_URL := os.Getenv("API_ORDER_URL")
	REDIS_URL := os.Getenv("REDIS_URL")
	DATABASE_URL := os.Getenv("DATABASE_URL")

	// Setup GRPC GESTION API
	serverGestion, err := grpc.Dial(API_GESTION_URL, grpc.WithInsecure())
	if err != nil { log.Fatalf("did not connect to API GESTION: %v", err) }
	defer serverGestion.Close()

	// Setup GRPC ORDER API
	serverOrder, err := grpc.Dial(API_ORDER_URL, grpc.WithInsecure())
	if err != nil { log.Fatalf("did not connect to API ORDER: %v", err) }
	defer serverOrder.Close()

	s := gateway.NewService(REDIS_URL, DATABASE_URL)

	g := gin.Default()

	s.SetupRoute(g, serverGestion, serverOrder)

	g.Run()
}
