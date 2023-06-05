package controllers

import "github.com/gin-gonic/gin"


func CreateQuizController(c *gin.Context){
	c.HTML(200,"createquiz.html",nil)
}