package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func GetAdminQuizController(c *gin.Context) {
	var quiz []models.Quiz
	dbconnection.DB.Find(&quiz)
	fmt.Println(quiz)
	c.HTML(200, "adminquizlist.html", gin.H{
		"quiz": quiz,
	})
}
