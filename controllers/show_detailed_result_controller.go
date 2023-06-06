package controllers

import (
	// "encoding/json"

	"encoding/json"
	"fmt"
	"reflect"

	// "reflect"

	// "reflect"
	// "strings"

	// "reflect"
	// "log"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowDetailedResultController(c *gin.Context) {
	fmt.Println("show detailed result controller called")
	session := sessions.Default(c)
	userid := session.Get("userID")
	quizname = c.Query("quizname")
	var questions map[string]interface{}

	dbconnection.DB.Model(&models.Quiz_Student{}).Select("result").Where("student_id=? and quiz_name=?", userid, quizname).Scan(&questions)

	fmt.Println(reflect.TypeOf(questions))
	result := questions["result"]

	resultbyte:=[]byte(result.(string))

	var results []models.Result

	err:=json.Unmarshal(resultbyte,&results)
	if err!=nil{
		fmt.Println(err.Error())
	}else{
		fmt.Println(results)
	}
	

	// ques := result.(string)
	// ques1:=strings.TrimPrefix(ques,"[")
	// ques2:=strings.TrimSuffix(ques1,"]")
	
	// // fmt.Println(ques)
	// // fmt.Println(reflect.TypeOf(ques))
	// // for i:=0;i<len(ques);i++{
	// // 	fmt.Print(string(ques[i]))
	// // }
	// // fmt.Println(ques2)
	// quesArr := strings.SplitN(ques2, ",",2)
	// fmt.Println(len(quesArr))

	// quesArr2:=strings.Join(quesArr,",")
	// quesArr3:=strings.Split(quesArr2,",")
	// fmt.Println(len(quesArr3))
	// fmt.Println(reflect.TypeOf(quesArr))
	// for i:=0;i<len(quesArr);i++{
	// 	fmt.Println(quesArr[0])
	// }
	c.HTML(200,"student_detailed_result.html",gin.H{
		"data":results,
	})
	// fmt.Println(quesArr)
	// var results []models.Result
	// for i := 0; i < len(quesArr); i++ {
	// 	jsonObj, _ := json.Marshal(quesArr[i])
	// 	// fmt.Println(string(jsonObj))
	// 	bytes:=[]byte(jsonObj)

	// 	var res models.Result

	// 	err:=json.Unmarshal(bytes,&res)
	// 	if err!=nil{
	// 		fmt.Println(err.Error())
	// 	}
	// 	fmt.Println(res)

	// }
	// fmt.Println(results)

}
