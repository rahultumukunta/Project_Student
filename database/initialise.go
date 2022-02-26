package database

import (
	"fmt"
	"log"
	"rahulProj/student/beans"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	// fmt.Println("in db connection func")

	dsn := "root:pass@tcp(127.0.0.1:3306)/projstudent?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("unable to connect to database")
	}

	DB = db

	student := beans.Student{
		Name:        "rahul",
		YearOfStudy: 1,
		Department:  "CSE",
		BloodGroup:  "A+",
		PhoneNo:     1234567891,
		Email:       "rahul@gmail.com",
		Location:    "Alwal",
	}

	db.AutoMigrate(&beans.Student{})
	if err := db.Create(&student).Error; err != nil {
		fmt.Println("db connection error2 ")
		log.Fatalln(err)
	}

}
