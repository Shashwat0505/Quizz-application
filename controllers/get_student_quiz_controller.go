package controllers

import (
	"encoding/json"
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)
var quizname string
func GetStudentQuizController(c *gin.Context) {
	quizname = c.Query("quizname")
	var questions []models.Questions
	dbconnection.DB.Debug().Model(&models.Quiz{}).Select("questions.question,questions.option_a,questions.option_b,questions.option_c,questions.option_d").Joins("inner join questions on questions.quiz_id=quizzes.id").Where("quizzes.quiz_name=?", quizname).Scan(&questions)
	fmt.Println(questions)
	j, _ := json.Marshal(questions)
	fmt.Println(string(j))

	//c.JSON(200, j)

	c.HTML(200, "studentquiz.html", gin.H{
		"questions": questions,
	})

}

func GetStudentQuizDataController(c *gin.Context) {
	
	var questions []models.Questions
	dbconnection.DB.Debug().Model(&models.Quiz{}).Select("questions.question,questions.option_a,questions.option_b,questions.option_c,questions.option_d").Joins("inner join questions on questions.quiz_id=quizzes.id").Where("quizzes.quiz_name=?", quizname).Scan(&questions)
	fmt.Println(questions)
	// j, _ := json.Marshal(questions)
	// fmt.Println(string(j))

	c.JSON(200, questions)

}


