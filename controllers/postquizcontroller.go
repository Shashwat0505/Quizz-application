package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	// "quizz-application/dbconnection"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func PostQuizController(c *gin.Context) {
	quizname :=c.Query("quizname")
	fmt.Println(quizname)
	fmt.Println("post request called")
	// fmt.Println(c.Request.Body)
	var questions []models.Questions

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithError(400, err)
		return
	}
	err = json.Unmarshal(body, &questions)
	fmt.Println(len(questions))
	if err != nil {
		c.AbortWithError(400, err)
		return

	}
	fmt.Println(quizname)
	fmt.Println(questions)
	q := models.Quiz{
		QuizName: quizname,
	}

	dbconnection.DB.Create(&q)

	for i := 0; i < len(questions); i++ {
		//
		questions[i].QuizId = q.ID
		dbconnection.DB.Create(&questions[i])
	}

	fmt.Println("data added!!!")

}
