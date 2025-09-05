package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/flights", createFlight)
	server.GET("/flights/:id", getFlightById)
}