package routes

import (
	"dunky.com/eventbooking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	// Route handlers

	server.GET("/api/v1/events", getEvents)     //GET, POST, PUT, PATCH, DELETE
	server.GET("/api/v1/events/:id", getEvent)  // api/v1/events/1, api/v1/events/2, etc
	authenticated := server.Group("/")          // To facilitate the implementation of Middleware
	authenticated.Use(middlewares.Authenticate) // To authenticate protected routes
	authenticated.POST("/api/v1/events", createEvent)
	authenticated.PUT("/api/v1/events/:id", updateEvent)
	authenticated.DELETE("/api/v1/events/:id", deleteEvent)
	authenticated.POST("/api/v1/events/:id/register", RegisterForEvent)
	authenticated.DELETE("/api/v1/events/:id/register", CancelRegistration)
	server.POST("/api/v1/signup", signup)
	server.POST("/api/v1/login", login)
}
