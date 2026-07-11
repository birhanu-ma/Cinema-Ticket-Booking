package routers

import (
	"cinema-ticket-booking/backend/services"

	"github.com/gin-gonic/gin"
)

func Routers(payment *services.Payment, user *services.Register) {
	routers := gin.Default()

	routers.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Cinema Ticket Booking API is running"})
	})

	routers.POST("/signup", user.Signup)
	routers.POST("/login", user.Login)

	// Hasura Action handlers — called by Hasura, not directly by the frontend
	routers.POST("/actions/initiate-payment", payment.InitiatePaymentAction)

	// Third-party webhook — called by Chapa directly
	routers.POST("/payment/webhook", payment.PaymentWebhook)

	routers.Run(":8081")
}