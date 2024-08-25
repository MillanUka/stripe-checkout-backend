package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go/v79"
	"millanuka.com/stripe-checkout/routes"
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	router := routes.InitRouter()

	stripe.Key = os.Getenv("SECRET_STRIPE_KEY")

	router.Run(os.Getenv("BASE_URL"))

}
