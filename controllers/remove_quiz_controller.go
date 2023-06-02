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
	dbconnection.DB.Delete(&models.Quiz{}).Where("q")

}
