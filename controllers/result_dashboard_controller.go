package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ResultDashboardController(c *gin.Context) {
	fmt.Println("result dashboard called")
	session := sessions.Default(c)
	studentID := session.Get("userID")

	// var count int64
	// dbconnection.DB.Model(&models.Quiz_Student{}).Where("student_id=?", studentID).Count(&count)
	// if count == 0 {
	// 	c.HTML(200, "studentpanel.html", gin.H{
	// 		"error": "You haven't attended any quiz!!",
	// 	})
	// 	return
	// }
	var result []models.Quiz_Student
	dbconnection.DB.Model(&models.Quiz_Student{}).Select("quiz_name,total_score").Where("student_id=?", studentID).Find(&result)
	fmt.Println(result)	
	c.HTML(200,"result_dashboard.html",gin.H{
		"quizzes":result,
	})

}
