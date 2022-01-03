package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `int:"age"`
}

var Users []User

func main() {
	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", UpdateUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func GetUsers(c *gin.Context) {
	c.JSON(200, Users)
}

func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	reqBody.ID = uuid.New().String()
	Users = append(Users, reqBody)

	c.JSON(200, gin.H{
		"error":   false,
		"message": "User created",
	})
}

func UpdateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age

			c.JSON(200, gin.H{
				"error":   false,
				"message": "User updated",
			})

			return
		}
	}

	c.JSON(422, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)

			c.JSON(200, gin.H{
				"error":   false,
				"message": "User deleted",
			})

			return
		}
	}

	c.JSON(422, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})
}
