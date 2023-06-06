package controllers

import (
	"fmt"
	"quizz-application/dbconnection"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetStudentResultController(c *gin.Context) {
	session := sessions.Default(c)
	teacherID := session.Get("userID")

	type temp struct {
		Name       string
		QuizName   string
		TotalScore string
	}
	var t1 []temp

	dbconnection.DB.Raw("select users.name,quiz_students.quiz_name,quiz_students.total_score from users left join quiz_students on users.id=quiz_students.student_id where users.id in(select student_id from teacher_students where teacher_students.teacher_id=?)", teacherID).Scan(&t1)

	fmt.Println(t1)
	c.HTML(200, "show_student_result.html", gin.H{
		"data": t1,
	})
}
