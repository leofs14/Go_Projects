package routes

import (
	"example.com/event-api/middlewares"
	"github.com/gin-gonic/gin"
)

//GET POST PUT PATCH DELETE
func RegisterRoutes(server *gin.Engine) { //get and create events can be used without import because are from the same package
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //dynamic path handler

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("/events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}