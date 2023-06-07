package controllers

import (
	"fmt"
	"net/http"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func RemoveUserController(c *gin.Context) {
	username := c.Query("username")
	fmt.Println(username)
	var userID int
	dbconnection.DB.Model(&models.User{}).Select("id").Where("name=?", username).Scan(&userID)
	dbconnection.DB.Where("id=?",userID).Delete(&models.User{})
	
	var users []models.User

	dbconnection.DB.Model(&models.User{}).Find(&users)

	// c.HTML(200, "deleteduser.html", gin.H{
	// 	"users": users,
	// 	"msg":     "user deleted successfully!!!!",
	// })


	c.Redirect(http.StatusTemporaryRedirect,"/admin/userprofile")
}
