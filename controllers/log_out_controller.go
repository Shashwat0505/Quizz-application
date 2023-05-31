package controllers

import "github.com/gin-gonic/gin"


func LogOutController(c *gin.Context){
	c.Redirect(301,"/authentication/Login")
	return
}