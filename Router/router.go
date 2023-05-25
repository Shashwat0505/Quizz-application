package Router

import (
	"fmt"
	"quizz-application/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Run() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("views/*")

	

	Authentication := r.Group("/authentication",ClearSessionHandler)
	{
		Authentication.GET("/Registration", controllers.RegistrationController)
		Authentication.POST("/PostRegistration", controllers.PostRegistrationDataController)
		Authentication.GET("/Login", controllers.LoginController)
		Authentication.POST("/Login", controllers.PostLoginDataController)
	}

	Admin := r.Group("/admin")
	{
		Admin.GET("/adminpanel", controllers.AdminpanelController)
		Admin.GET("/addteacher", controllers.AddTeacherController)
		Admin.GET("/addadmin", controllers.AddAdminController)
		Admin.POST("/addteacher", controllers.CreateTeacherController)
		Admin.POST("/addadmin", controllers.CreateAdminController)

	}

	Teacher := r.Group("/teacher", SessionHandler)
	{
		Teacher.GET("/teacherpanel", controllers.TeacherPanelController)
		Teacher.GET("/createquizz", controllers.CreateQuizController)
		Teacher.POST("/createquizz",controllers.PostQuizController)
	}

	r.Run(":8080")

}

func SessionHandler(c *gin.Context) {
	fmt.Println("session is called")
	
	session := sessions.Default(c)
	fmt.Println(session.Get("userID"))
	if session.Get("userID") == nil {
		c.HTML(200, "Login.html", nil)
		return
	}
}

func ClearSessionHandler(c *gin.Context){
 session:=sessions.Default(c)
 session.Clear()
 session.Save()
 
}