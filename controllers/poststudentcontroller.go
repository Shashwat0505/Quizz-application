package controllers

import (
	"fmt"
	"log"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostStudentController(c *gin.Context) {
	fmt.Println("post student controller called")
	name:=c.Request.FormValue("studentname")
	fmt.Println(name)
	err := c.Bind(&name)
	if err != nil {
		log.Fatal(err)
	}
	if name == "" {
		c.HTML(400, "addstudent.html", gin.H{
			"error": "Please enter valid input!!!!",
		})
		return
	}
	var count int64
	var user models.User
	dbconnection.DB.Where("name=?", name).Find(&user).Count(&count)
	if count == 0 {
		c.HTML(400, "addstudent.html", gin.H{
			"error": "User not found!!!!",
		})
		return
	}
	session := sessions.Default(c)
	var teacherid = session.Get("userID").(int)
	var student models.User
	dbconnection.DB.Model(&models.User{}).Select("id").Where("name=?", name).Find(&student)
	studentid := student.ID
	dbconnection.DB.Create(&models.Teacher_Student{TeacherID: teacherid, StudentID: studentid})

	c.HTML(200, "addstudent.html", gin.H{
		"error": "data added!!!",
	})

}