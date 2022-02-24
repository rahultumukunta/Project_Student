package main

import (
	"rahulProj/student/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controllers.StudentDetails)
	router.GET("/students/:id", controllers.StudentByID)
	router.POST("/register", controllers.RegisterStudent)

	router.Run(":8080")
}
