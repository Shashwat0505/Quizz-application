package controllers

import "github.com/gin-gonic/gin"


func LoginController(c *gin.Context){
	c.HTML(200,"login.html",nil)
}