package main

import (
	"rahulProj/student/controllers"
	"rahulProj/student/database"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/", controllers.StudentDetails)
	router.GET("/students/:id", controllers.StudentByID)
	router.POST("/register", controllers.RegisterStudent)
	router.GET("./unregister/:id", controllers.UnregisterStudent)

	database.DatabaseInit()

	router.Run(":8080")
}
