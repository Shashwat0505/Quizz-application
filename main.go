package main

import (
	"quizz-application/Router"
	"quizz-application/dbconnection"
)

func init() {
	dbconnection.DBconnection()
}

func main() {
	Router.Run()
}
