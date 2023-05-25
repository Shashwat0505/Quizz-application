package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AddAdminController(c *gin.Context){
	session:=sessions.Default(c)
	if session.Get("userID")!=nil{
	c.HTML(200,"addadmin.html",nil)
	}else{
		c.HTML(200,"login.html",nil)
	}
}