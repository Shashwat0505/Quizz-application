package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func PostQuizController(c *gin.Context) {
	fmt.Println("post request called")
	var question map[string]interface{}

	c.ShouldBind(&question)
	fmt.Println(question)

}
