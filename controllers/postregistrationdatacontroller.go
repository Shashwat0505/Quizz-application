package controllers

import (
	"fmt"
	"log"
	"quizz-application/dbconnection"
	"quizz-application/models"

	"github.com/gin-gonic/gin"
)

func PostRegistrationDataController(c *gin.Context) {
	var user models.User

	c.ShouldBind(&user)


	if user.Name==""||user.Email==""||user.Password==""{
		c.HTML(301, "registration.html",gin.H{
			"error":"can't enter empty field!!!",
		})
		return
	}
	fmt.Println(user)
	var count int64

	dbconnection.DB.Model(&models.User{}).Where("email=? or password=?", user.Email, user.Name).Count(&count)
	fmt.Println(count)

	if count > 0 {
		c.HTML(301, "registration.html",gin.H{
			"error":"email or password is already taken!!!!",
		})
		return

	}

	if user.Password == "admin@123" {
		newuser := models.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
			RoleName: "admin",
			Role: models.Role{
				Role_name: "admin",
			},
			
		}
		if err := dbconnection.DB.Model(&models.User{}).Create(&newuser).Error; err != nil {
			log.Fatal(err.Error)
		}

		n := fmt.Sprintf("welcome %s !!!", newuser.Name)
		c.String(200, n)
		return

	}

	newuser := models.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		RoleName: "student",
		Role: models.Role{
			Role_name: "student",
		},
	}
	if err := dbconnection.DB.Model(&models.User{}).Create(&newuser).Error; err != nil {
		log.Fatal(err.Error)
	}
	c.HTML(200,"login.html",nil)

}
