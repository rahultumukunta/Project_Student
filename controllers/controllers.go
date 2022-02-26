package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rahulProj/student/beans"
	"rahulProj/student/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StudentDetails(c *gin.Context) {

	var StudentInfo []beans.Student

	database.DB.Find(&StudentInfo)
	// fmt.Println("Get All rows from db")
	// fmt.Println(StudentInfo)
	c.IndentedJSON(200, StudentInfo)
}

func StudentByID(c *gin.Context) {
	id := c.Param("id")

	id_int, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		fmt.Printf("Eroor while converting id string to int")
	}

	// er := 0
	var StudentInfo beans.Student

	result := database.DB.Where("id = ?", id_int).First(&StudentInfo)
	if result.Error != nil {
		fmt.Println("error getting value by ID")
		fmt.Println(result.Error)
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"Message": " cannot find student by ID " + id,
		})
	} else {
		c.IndentedJSON(200, StudentInfo)
	}
	// fmt.Println("get Row By Id ")
	// fmt.Println(StudentInfo)

}

func RegisterStudent(c *gin.Context) {

	var StudentInfo beans.Student
	// er := 0

	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(requestBody, &StudentInfo)

	result := database.DB.Create(&StudentInfo)
	if result.Error != nil {
		fmt.Println("db connection error while registering student ")
		fmt.Println(result.Error)
		c.IndentedJSON(400, "Registration unsuccesful")
		// er = 1
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": StudentInfo.Name + " Registered Succesfully",
		})
	}

	// fmt.Printf("%+v\n", StudentInfo)
	// tempdetails.StudentInfo = append(tempdetails.StudentInfo, beans.StudentBean(StudentInfo))

}

func UnregisterStudent(c *gin.Context) {

	id := c.Param("id")

	id_int, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		fmt.Printf("Eroor while converting id string to int")
	}

	var StudentInfo beans.Student

	result := database.DB.Where("id = ?", id_int).Delete(&StudentInfo)
	if result.Error != nil {
		fmt.Println("Error while unregistering student")
		fmt.Println(result.Error)
		c.IndentedJSON(400, "Unregistration unsuccesful")
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": StudentInfo.Name + " Unregistered Succesfully",
		})
	}

}
