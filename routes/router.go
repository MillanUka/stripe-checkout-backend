package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"millanuka.com/stripe-checkout/controllers"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/create-checkout-session", controllers.CreateCheckoutSession)

	return router
}
