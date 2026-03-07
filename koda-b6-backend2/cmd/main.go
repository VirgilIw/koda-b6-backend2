package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/virgilIw/koda-b6-backend2/internal/config"
	"github.com/virgilIw/koda-b6-backend2/internal/di"
	"github.com/virgilIw/koda-b6-backend2/internal/router"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	db, err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rdb := config.InitRedis()

	app := gin.Default()

	container := di.NewContainer(db, rdb)

	router.InitRouter(app, container)

	port := os.Getenv("PORT")

	app.Run(fmt.Sprintf("localhost:%s", port))
}
