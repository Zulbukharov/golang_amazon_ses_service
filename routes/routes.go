package routes

import (
	"fmt"

	"github.com/Zulbukharov/golang_amazon_ses_service/handlers"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes ...
func InitializeRoutes(router *gin.Engine) {
	router.GET("/", handlers.Cors(), handlers.Hello)
	router.GET("/email", handlers.Cors(), handlers.Email)
	fmt.Println("routes initialized")
}
