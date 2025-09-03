package routes

import (
	"net/http"

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