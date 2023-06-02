package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func RemoveQuizController(c *gin.Context) {
	quizname := c.Query("quizname")
	fmt.Println(quizname)
	var quizid int
	dbconnection.DB.Model(&models.Quiz{}).Select("id").Where("quiz_name=?",quizname).Scan(&quizid)
	dbconnection.DB.Where("quiz_id=?",quizid).Delete(&models.Questions{})
	dbconnection.DB.Where("id=?",quizid).Delete(&models.Quiz{})		
	var quizzes []models.Quiz

	dbconnection.DB.Model(&models.Quiz{}).Find(&quizzes)

	c.HTML(200,"deletedquiz.html",gin.H{
		"quizzes":quizzes,
		"msg":"quiz deleted successfully!!!!",
	})
	


}
