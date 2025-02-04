package routes

import (
	"net/http"
	"restapi/middleware"
	"restapi/models"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", func(context *gin.Context) {
		events, err := models.GetAllEvent()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not fetch data , try again latter"})
		}
		context.JSON(http.StatusOK, events)
	})

	server.GET("/events/:id", getEventById) // event/1 or event/5
    authenicated:=server.Group("/")  // Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix. 
    authenicated.Use(middleware.Authenticate)
	authenicated.POST("/events",eventsCreated)
    authenicated.PUT("/events/:id",updateEvent)
	authenicated.DELETE("/events/:id",deleteEvent)
	authenicated.POST("/events/:id/register",registerForEvent)
	authenicated.DELETE("/events/:id/register",cancelRegistration)

	// server.POST("/events",middleware.Authenticate, eventsCreated)   // create event i.e write to server , this is also wasy to add middleware  
	// server.PUT("/events/:id", updateEvent)
	// server.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", signup)
	server.GET("/signup/:id", getuserById)
	server.POST("/login", login)
	server.Run(":8081")

}
