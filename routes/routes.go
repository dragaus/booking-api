package routes

import (
	"example.com/booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	//User endpoints
	server.POST("/signup", signup)
	server.POST("/login", login)

	// Events endpoints
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	// Authenticated routes
	authenticated := server.Group("/", middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)

}
