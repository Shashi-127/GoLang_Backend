package routes

import (
	"net/http"
	"restapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param returns the value of the URL param
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't fetch Event"})
	}
	err = event.Register(userId)
	if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"couldn't Register Event"})
	}
	context.JSON(http.StatusCreated,gin.H{"message":"registered "})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param returns the value of the URL param
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	var event models.Event 
	event.Id=eventId
	err=event.CancelRegistration(userId)
    if err!=nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message":"couldn't Cancel Event"})
	}
	context.JSON(http.StatusOK,gin.H{"message":"Cancelled"})
}
