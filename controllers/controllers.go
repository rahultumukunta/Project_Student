package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"rahulProj/student/beans"
	tempdetails "rahulProj/student/tempDetails"
	"strconv"

	"github.com/gin-gonic/gin"
)

func StudentDetails(c *gin.Context) {
	c.IndentedJSON(200, tempdetails.StudentInfo)
}

func StudentByID(c *gin.Context) {
	id := c.Param("id")

	id_int, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		fmt.Printf("Eroor while converting id string to int")
	}

	for _, val := range tempdetails.StudentInfo {
		if val.ID == int(id_int) {
			c.IndentedJSON(200, val)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{
		"Message": " cannot find student by ID " + id,
	})
}

func RegisterStudent(c *gin.Context) {

	var StudentInfo beans.StudentBean

	requestBody, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(requestBody, &StudentInfo)

	c.JSON(http.StatusOK, gin.H{
		"message": StudentInfo.Name + " Registered Succesfully",
	})

	fmt.Printf("%+v\n", StudentInfo)
	tempdetails.StudentInfo = append(tempdetails.StudentInfo, beans.StudentBean(StudentInfo))

}
