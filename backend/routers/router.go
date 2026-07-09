package routers

import (
	"cinema-ticket-booking/backend/services"

	"github.com/gin-gonic/gin"
)



func Routers( user *services.Register){
	routers:=gin.Default()
	routers.GET("/", func(c *gin.Context) {
    c.JSON(200, gin.H{
        "message": "Cinema Ticket Booking API is running",
    })
})
    routers.POST("/signup", user.Signup)
	routers.POST("/login", user.Login)

	routers.Run(":8081")


}