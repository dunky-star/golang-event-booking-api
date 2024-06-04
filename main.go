package main

import (
	"net/http"

	"dunky.com/eventbooking/models"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/events", createEvent)

	server.Run(":9090") //localhost:9090

} // End of main()

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request body."})
		return
	}
	event.ID = 1
	event.UserId = 1
	event.Save() // Saving the event model
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}
