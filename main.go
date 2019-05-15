package main

import (
	"log"
	"os"

	"github.com/Zulbukharov/golang_amazon_ses_service/db"
	"github.com/Zulbukharov/golang_amazon_ses_service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	routes.InitializeRoutes(router)
	db.Init()
	router.Run(":" + port)
}
