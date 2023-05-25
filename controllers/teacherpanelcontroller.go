package controllers

import "github.com/gin-gonic/gin"

func TeacherPanelController(c *gin.Context){
	c.HTML(200,"teacherpanel.html",nil)
}