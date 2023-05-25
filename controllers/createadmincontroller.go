package controllers

import (
	"fmt"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func CreateAdminController(c *gin.Context) {
	
	name := c.PostForm("adminname")
	fmt.Println(name)

	var newuser models.User
	var count int64
	dbconnection.DB.Debug().Where("name=?", name).Find(&newuser).Count(&count)
	fmt.Println(count)
	if count == 0 {
		c.HTML(404, "addteacher.html", gin.H{
			"error": "user not found!!!",
		})
		return
	}

	dbconnection.DB.Debug().Model(&models.User{}).Where("name=?", name).Update("role_name", "admin")
	c.HTML(200, "addteacher.html", gin.H{
		"error": "Role changed successfully!!!",
	})
}
