package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func GetQuizController(c *gin.Context) {

	fmt.Println("get quiz controller called")
	quizname := c.Query("quizname")
	var questions []models.Questions
	var quizzes []models.Quiz
	dbconnection.DB.Debug().Model(&models.Quiz{}).Select("quizzes.quiz_name,questions.question,questions.option_a,questions.option_b,questions.option_c,questions.option_d,questions.answer,questions.difficulty").Joins("inner join questions on questions.quiz_id=quizzes.id").Where("quizzes.quiz_name=?", quizname).Scan(&questions)
	dbconnection.DB.Find(&quizzes)
	fmt.Println(questions)
	c.HTML(200,"showquiz.html",gin.H{
		"questions":questions,
		"name":quizname,
		"quizzes":quizzes,
	})

	
	
	
	
}