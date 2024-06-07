package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	// Route handlers
	server.GET("/api/v1/events", getEvents)    //GET, POST, PUT, PATCH, DELETE
	server.GET("/api/v1/events/:id", getEvent) // api/v1/events/1, api/v1/events/2, etc
	server.POST("/api/v1/events", createEvent)
	server.POST("/api/v1/signup", signup)
	server.POST("/api/v1/login", login)
}
