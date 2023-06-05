package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func GetUserAdminController(c *gin.Context) {
	var users []models.User
	dbconnection.DB.Find(&users)
	fmt.Println(users)
	c.HTML(200, "adminuserlist.html", gin.H{
		"users": users,
	})
}
