package routers

import (
	"cinema-ticket-booking/backend/services"

	"github.com/gin-gonic/gin"
)

func Routers(payment *services.Payment, user *services.Register, upload *services.Upload) {
	routers := gin.Default()

	routers.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Cinema Ticket Booking API is running"})
	})

	routers.POST("/signup", user.Signup)
	routers.POST("/login", user.Login)
	routers.Static("/uploads", upload.UploadDir)

	routers.POST("/actions/create-booking", payment.CreateBookingAction)
	routers.POST("/actions/initiate-payment", payment.InitiatePaymentAction)
	routers.POST("/actions/upload-image", upload.UploadImageAction)

	routers.Any("/payment/webhook", payment.PaymentWebhook)

	routers.Run(":8081")
}