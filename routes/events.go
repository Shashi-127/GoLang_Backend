package routes

import (
	"net/http"
	"restapi/models"

	"strconv"

	"github.com/gin-gonic/gin"
)

func getEventById(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param returns the value of the URL param
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	context.JSON(http.StatusOK, event)
}
func eventsCreated(context *gin.Context) {
	
	var event models.Event
	err:= context.ShouldBindJSON(&event) // attempts to parse the incoming JSON payload from the request body and bind it to the event struct.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed requested data"})
		return
	}
	userId:=context.GetInt64("userId")
	event.UserId =userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch data , try again latter"})
		return

	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event Created", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param returns the value of the URL param
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId:=context.GetInt64("userId")
    event, err:= models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch  event"})
		return
	}
	if event.UserId!=userId {
       context.JSON(http.StatusUnauthorized,gin.H{"message":"Not authorized to update Event"})
	   return 
	}
	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent) // attempts to parse the incoming JSON payload from the request body and bind it to the event struct.
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parsed requested data"})
		return
	}
	updateEvent.Id = eventId
	err = updateEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event updated successfully!"})

}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64) // Param returns the value of the URL param
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	userId:=context.GetInt64("userId")
	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch  event"})
		return
	}
	if event.UserId!=userId {
		context.JSON(http.StatusUnauthorized,gin.H{"message":"Not authorized to Delete Event"})
		return 
	 }

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not Delete event"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Event Successfully deleted!!"})
}
