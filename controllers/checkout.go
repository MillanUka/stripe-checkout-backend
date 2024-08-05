package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v79"
	"github.com/stripe/stripe-go/v79/checkout/session"
	"log"
	"net/http"
)

func CreateCheckoutSession(context *gin.Context) {
	domain := "http://localhost:8001"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			{
				Price:    stripe.String("{{PRICE_ID}}"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "?success=true"),
		CancelURL:  stripe.String(domain + "?canceled=true"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	context.Redirect(http.StatusSeeOther, s.URL)
}
