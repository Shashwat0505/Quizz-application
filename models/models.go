package models

type User struct {
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	RoleName string `json:"rolename" form:"rolename"`
	Role     Role   `gorm:"foreignKey:RoleName;References:Role_name"`
}
type Role struct {
	Role_name string `gorm:"primaryKey"`

}

type Questions struct{
	ID int `gorm:"primaryKey"`
	Question string `json:"question" form:"question"`
	OptionA  string `json:"optiona" form:"optiona"`
	OptionB  string  `json:"optionb" form:"optionb"`
	OptionC  string  `json:"optionc" form:"optionc"`
	OptionD  string   `json:"optiond" form:"optiond"`
	Answer   string   `json:"answer" form:"answer"`
	Difficulty string  `json:"difficulty" form:"difficulty"`
	QuizId int         `json:"quizid" form:"quizid"`
	Quiz Quiz `gorm:"foreignKey:QuizId;References:ID"`
}

type Quiz struct{
	ID int `gorm:"primaryKey"`
	QuizName string `json:"quizname" form:"quizname"`
	CreatorID int


}