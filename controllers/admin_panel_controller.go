package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AdminpanelController(c *gin.Context) {
	session := sessions.Default(c)
	if session.Get("userID") != nil {
		c.HTML(200, "adminpanel.html", nil)
	}else{
		c.HTML(200,"login.html",nil)
	}
}
