package controllers

import (
	"encoding/json"
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SubmitQuizController(c *gin.Context) {
	session:=sessions.Default(c)

	type TempStruct struct {
		Answer   string `json:"answer"`
		Question string `json:"question"`
		Selected string `json:"selected"`
		Status   string `json:"status"`
	}

	type QuizStruct struct {
		QuizName   string       `json:"quizname"`
		Result     []TempStruct `json:"result"`
		TotalScore string  	        `json:"total_score"`
	}
	var quiz_struct QuizStruct
	var quiz_student models.Quiz_Student
	fmt.Println("submit controller called")
	// var m map[string]interface{}
	err := c.ShouldBindJSON(&quiz_struct)
	fmt.Println(err)
	fmt.Println("Hello", quiz_struct)
	// fmt.Println(quiz_struct)
	//

	// quiz_student.QuizName = m["quizname"].(string)
	// var res map[string]interface{}

	// var tempstruct []TempStruct
	// data, _ := json.Marshal(m["result"])
	// json.Unmarshal(data, &tempstruct)
	// fmt.Println(tempstruct)
	
		sliceData, _ := json.Marshal(quiz_struct.Result)
		quiz_student.Result = string(sliceData)
	
	// quiz_student.Result = strings.Split(string(data), ",")
	// fmt.Println(quiz_student.Result)
	// temp := strings.Split(quiz_student.Result[0], "[")
	// quiz_student.Result[0] = strings.TrimSpace(temp[0])
	// temp = strings.Split(quiz_student.Result[len(quiz_student.Result)-1], "]")
	// quiz_student.Result[len(quiz_student.Result)-1] = strings.TrimSpace(temp[0])
	fmt.Println(quiz_student.Result)
	quiz_student.QuizName=quiz_struct.QuizName
	quiz_student.StudentID=session.Get("userID").(int)
	quiz_student.TotalScore=quiz_struct.TotalScore

	dbconnection.DB.Create(&quiz_student)
	fmt.Println("data added")	



}
