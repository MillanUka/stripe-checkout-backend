package routes

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"millanuka.com/stripe-checkout/controllers"
	"millanuka.com/stripe-checkout/db"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(cors.Default())
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.POST("/create-checkout-session", controllers.CreateCheckoutSession)
	router.GET("/products", func(ctx *gin.Context) {
		products, err := db.LoadAllProducts()

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": "error"})
		}
		ctx.JSON(http.StatusOK, products)
	})

	return router
}
