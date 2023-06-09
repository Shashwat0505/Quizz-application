package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func ShowStudentProfileController(c *gin.Context) {
	session := sessions.Default(c)
	userid := session.Get("userID")

	

	type TempStudent struct {
		StudentName      string
		QuizzesAttended  []string
		QuizzesRemaining []string
	}
	var tempStudentArr []TempStudent
	var students []string

	dbconnection.DB.Model(&models.User{}).Select("users.name").Joins("inner join teacher_students on teacher_students.student_id=users.id").Where("teacher_students.teacher_id=?", userid).Scan(&students)
	
	for i:=0;i<len(students);i++{
		var studentid int
		dbconnection.DB.Model(&models.User{}).Select("id").Where("name=?",students[i]).Scan(&studentid)
		var quizzesparticipated []string
		dbconnection.DB.Model(&models.Quiz_Student{}).Select("quiz_name").Where("student_id=?",studentid).Scan(&quizzesparticipated)

		var teacherquizzes []string
		dbconnection.DB.Model(&models.Quiz{}).Select("quiz_name").Where("creator_id=?",userid).Scan(&teacherquizzes)
	

		var remainingQuizzes []string
		remainingQuizzes=difference(quizzesparticipated,teacherquizzes)
		
		t:=TempStudent{
			StudentName :students[i],
			QuizzesAttended:quizzesparticipated,
			QuizzesRemaining:remainingQuizzes,
		}
		tempStudentArr=append(tempStudentArr, t)

		
	}
	fmt.Println(tempStudentArr)

	c.HTML(200,"studentprofiles.html",gin.H{
		"studentprofiles":tempStudentArr,
	})
	return
}


func difference(slice1 []string, slice2 []string) []string {
    var diff []string

    // Loop two times, first to find slice1 strings not in slice2,
    // second loop to find slice2 strings not in slice1
    for i := 0; i < 2; i++ {
        for _, s1 := range slice1 {
            found := false
            for _, s2 := range slice2 {
                if s1 == s2 {
                    found = true
                    break
                }
            }
            // String not found. We add it to return slice
            if !found {
                diff = append(diff, s1)
            }
        }
        // Swap the slices, only if it was the first loop
        if i == 0 {
            slice1, slice2 = slice2, slice1
        }
    }

    return diff
}