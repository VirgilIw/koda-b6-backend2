package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/virgilIw/koda-b6-backend2/internal/di"
	"github.com/virgilIw/koda-b6-backend2/internal/router"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	}

	r := gin.Default()

	container := di.NewContainer()
	router.Init(r, container)

	port := os.Getenv("PORT")

	r.Run(fmt.Sprintf("localhost:%s", port))
}
