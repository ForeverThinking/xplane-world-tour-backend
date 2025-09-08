package routes

import (
	"net/http"
	"strconv"

	"github.com/ForeverThinking/xplane-world-tour-backend/models"
	"github.com/gin-gonic/gin"
)

func createFlight(context *gin.Context) {
	var flight models.Flight

	if err := context.ShouldBindJSON(&flight); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request,"})
		return
	}

	if err := flight.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save flight."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Flight created.", "flight": flight})
}

func getAllFlights(context *gin.Context) {
	flights, err := models.GetAllFlights()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get flights."})
		return
	}

	context.JSON(http.StatusOK, flights)
}

func getFlightById(context *gin.Context) {
	flightId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse input."})
		return
	}

	flight, err := models.GetFlightById(flightId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find flight."})
		return
	}

	context.JSON(http.StatusOK, flight)
}
