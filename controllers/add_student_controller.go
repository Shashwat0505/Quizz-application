package controllers

import "github.com/gin-gonic/gin"

func AddStudentController(c *gin.Context){
	c.HTML(200,"addstudent.html",nil)
}