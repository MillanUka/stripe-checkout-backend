package main

import (
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"log"
	"millanuka.com/stripe-checkout/routes"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	stripe.Key = os.Getenv("SECRET_STRIPE_KEY")

	router := routes.InitRouter()

	router.Run("localhost:8001")
}
