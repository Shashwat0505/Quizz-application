package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func UpdateQuizController(c *gin.Context) {

	quizname := c.Query("quizname")
	var quizid int
	dbconnection.DB.Model(&models.Quiz{}).Select("id").Where("quiz_name=?", quizname).Scan(&quizid)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	var questions []models.Questions
	err = json.Unmarshal(body, &questions)
	fmt.Println(len(questions))
	if err != nil {
		c.AbortWithError(400, err)
		return

	}
	for i := 0; i < len(questions); i++ {
		//
		questions[i].QuizId = quizid
		dbconnection.DB.Create(&questions[i])
	}
}
