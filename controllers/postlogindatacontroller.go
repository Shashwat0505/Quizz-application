package controllers

import (
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func PostLoginDataController(c *gin.Context) {

	session:=sessions.Default(c)
	session.Clear()
	session.Save()
	var user models.User

	c.ShouldBind(&user)

	var newuser models.User
	if user.Email == "" || user.Password == "" {
		c.HTML(301, "login.html", gin.H{
			"error": "email or password is empty",
		})
		return
	}
	var count int64
	dbconnection.DB.Where("email=? and password=?", user.Email, user.Password).Find(&newuser).Count(&count)
	if count == 0 {
		c.HTML(301, "registration.html", gin.H{
			"error": "user does not exist!!",
		})
		return
	}
	if user.Password == "admin@123" && newuser.RoleName == "admin" {
		session.Set("userID",newuser.ID)
		session.Save()
		c.HTML(200, "adminpanel.html", gin.H{
			"name": newuser.Name,
		})
		return
	}

	if newuser.RoleName == "teacher" {
		session.Set("userID",newuser.ID)
		session.Save()
		c.HTML(200, "teacherpanel.html", gin.H{
			"name": newuser.Name,
		})
		return
	}

	if newuser.RoleName=="student"{
		session.Set("userID",newuser.ID)
		session.Save()
		c.Redirect(301,"/student/studentpanel")
		return

	}

}
