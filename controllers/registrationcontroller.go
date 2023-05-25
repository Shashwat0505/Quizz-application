package controllers

import (
	"github.com/gin-gonic/gin"

)
func RegistrationController( c *gin.Context){
	c.HTML(200,"registration.html",nil)
}