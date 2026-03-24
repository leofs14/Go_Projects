package routes

import "github.com/gin-gonic/gin"

//GET POST PUT PATCH DELETE
func RegisterRoutes(server *gin.Engine) { //get and create events can be used without import because are from the same package
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) //dynamic path handler
	server.POST("/events", createEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signup)
	server.POST("/login", login)
}