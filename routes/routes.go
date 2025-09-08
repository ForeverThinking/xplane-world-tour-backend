package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.POST("/flights", createFlight)
	server.GET("flights", getAllFlights)
	server.GET("/flights/:id", getFlightById)
	server.PUT("/flights/:id", updateFlightById)
	server.DELETE("/flights/:id", deleteFlight)
}
