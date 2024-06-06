package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Route handlers
	server.GET("/api/v1/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.POST("/api/v1/events", createEvent)
	server.POST("/api/v1/signup", signup)

}
