package main

import (
	"log"

	restHandler "api-tutorial/transport/rest/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", restHandler.GetUsers)
		userRoutes.POST("/", restHandler.CreateUser)
		userRoutes.PUT("/:id", restHandler.UpdateUser)
		userRoutes.DELETE("/:id", restHandler.DeleteUser)
	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}
