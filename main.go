package main

import (
	"restapi/routes"

	"restapi/db"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)

}

// package main

// Router:

// Matches incoming HTTP requests to the correct handler.
// Example: GET /about -> aboutHandler.
// Middleware:

// Intercepts and processes requests before they reach the handler.
// Useful for logging, authentication, error handling, etc.

// import (

// 	"github.com/gin-gonic/gin"
// )

// func main() {

// 	r := gin.Default() // r is router which will  It listens for incoming HTTP requests (like "GET /home" or "POST /login") and sends them to the correct handler function that processes those requests

// 	// func(c *gin.Context) { c.JSON(200, gin.H{"Name": "Shashi Kumar"}) }
// 	// This is the handler function, which defines what happens when the router matches the incoming request.
// 	// The handler generates the response by sending a JSON object with the data {"Name": "Shashi Kumar"} and an HTTP status code 200 (OK).

// 	r.GET("/events", func(c *gin.Context) { // r.GET("/") , this is route definition and part of the router
// 		c.JSON(200, gin.H{ // H is map of type map[string]any
// 			"Name": "Shashi Kumar", // key:value pair
// 		})
// 	})
// 	r.Run(":8081") // localhost:8080

// }

//Endpoints: Endpoints are the access points for clients to interact with your API.e.g GET/users
// They are defined by a path (e.g., /users) and an HTTP method (e.g., GET, POST).
// Router: Determines which handler should process an incoming request based on the HTTP method and URL path.
// Handler: A function that contains the logic to process the request and send the response.
// Middleware: A function or set of functions that intercept requests before they reach the handler to perform tasks like logging, authentication, etc
// ;
