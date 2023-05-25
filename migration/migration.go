package main

import (
	"quizz-application/dbconnection"
	"quizz-application/models"
)

func init() {
	dbconnection.DBconnection()
}
func main() {
	
	// dbconnection.DB.AutoMigrate(&models.User{}, &models.Role{})
	dbconnection.DB.AutoMigrate(&models.Questions{},&models.Quiz{})
}
