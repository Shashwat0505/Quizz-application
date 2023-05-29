package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func StudentPanelController(c *gin.Context) {

	session := sessions.Default(c)
	userid := session.Get("userID")
	var user models.User
	dbconnection.DB.Where("id=?", userid).Find(&user)
	if user.RoleName != "student" {
		c.HTML(200, "login.html", gin.H{
			"error": "unauthorized access",
		})
		return
	}
	var quizzes []models.Quiz

	dbconnection.DB.Model(&models.Teacher_Student{}).Select("quizzes.quiz_name").Joins("inner join quizzes on quizzes.creator_id=teacher_students.teacher_id ").Where("teacher_students.student_id=?", userid).Scan(&quizzes)
	fmt.Println(quizzes)

	c.HTML(200, "studentpanel.html", gin.H{
		"data":quizzes,
	})
}