package main

import (
	"backend/todo"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s := todo.NewService(os.Getenv("REDIS_URI"), os.Getenv("POSTGRESQL_URI"))

	g := gin.Default()
	s.SetupRoute(g)

	port := os.Getenv("PORT")
	g.Run(fmt.Sprint(":", port))
}
