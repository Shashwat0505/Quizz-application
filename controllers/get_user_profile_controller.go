package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func GetUserProfileController(c *gin.Context) {
	type TempStudent struct{
		Name string
		Email string
		RoleName string
		Quizzes []string
		Teacher string
	}
	type TempTeacher struct{
		Name string
		Email string
		RoleName string
		QuizzesCreated []string
		
	}
	
	var tempStudentSlice []TempStudent
	var tempTeacherSlice []TempTeacher
	var users []models.User

	dbconnection.DB.Find(&users)

	for i := 0; i < len(users); i++ {
		if users[i].RoleName == "student" {
			var quizzes []string
			dbconnection.DB.Model(&models.Quiz_Student{}).Select("quiz_name").Where("student_id=?", users[i].ID).Scan(&quizzes)
			fmt.Println(users[i].Name, quizzes)

			var teacherID int
			var teacher string
			dbconnection.DB.Model(models.Teacher_Student{}).Select("teacher_students.teacher_id").Where("teacher_students.student_id=?", users[i].ID).Scan(&teacherID)
			dbconnection.DB.Model(&models.User{}).Select("users.name").Where("users.id=?", teacherID).Scan(&teacher)
			fmt.Println(users[i].Name, teacher)

			
			
			t:=TempStudent{
				Name:users[i].Name,
				Email:users[i].Email,
				RoleName:users[i].RoleName,
				Quizzes:quizzes,
				Teacher:teacher,


			}
			// fmt.Println(t)
			tempStudentSlice=append(tempStudentSlice, t)
			


			
		}else if users[i].RoleName=="teacher" {
			var quizzes []string
			dbconnection.DB.Model(&models.Quiz{}).Select("quiz_name").Where("creator_id=?",users[i].ID).Scan(&quizzes)
			t2:=TempTeacher{
				Name: users[i].Name,
				Email: users[i].Email,
				RoleName: users[i].RoleName,
				QuizzesCreated: quizzes,

			}
			tempTeacherSlice=append(tempTeacherSlice, t2)
			
		}

		
	}
	fmt.Println(tempStudentSlice)
	fmt.Println(tempTeacherSlice)
	c.HTML(200,"userprofile.html",gin.H{
		"student":tempStudentSlice,
		"teacher":tempTeacherSlice,
	})

}
