package Router

import (
	"fmt"
	"quizz-application/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	// "github.com/itsjamie/gin-cors"
)

func Run() {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))
	r.LoadHTMLGlob("views/*")

	Authentication := r.Group("/authentication", ClearSessionHandler)
	{
		Authentication.GET("/Registration", controllers.RegistrationController)
		Authentication.POST("/PostRegistration", controllers.PostRegistrationDataController)
		Authentication.GET("/Login", controllers.LoginController)
		Authentication.POST("/Login", controllers.PostLoginDataController)
		Authentication.GET("/Logout",controllers.LogOutController)
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
		Teacher.POST("/createquizz", controllers.PostQuizController)
		Teacher.GET("/addstudent", controllers.AddStudentController)
		Teacher.POST("/addstudent", controllers.PostStudentController)
		Teacher.GET("/listofquiz", controllers.GetListOfQuizController)
		Teacher.GET("/quiz", controllers.GetQuizController)
		Teacher.GET("/showstudentresult",controllers.GetStudentResultController)
	}
	Student := r.Group("/student")
	{
		Student.GET("/studentpanel", controllers.StudentPanelController)
		Student.GET("/quiz", controllers.GetStudentQuizController)
		Student.GET("/quiz-data", controllers.GetStudentQuizDataController)
		Student.POST("/submit",controllers.SubmitQuizController)
		Student.GET("/resultdashboard",controllers.ResultDashboardController)
	}
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	r.Use(cors.New(cors.Config{
		//AllowOrigins:     []string{"*"},
		AllowAllOrigins:        true,
		AllowMethods:           []string{"PATCH", "GET", "POST", "OPTIONS"},
		AllowHeaders:           []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:          []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		
	}))
	// r.Use(cors.Default())
	// r.Use(cors.Middleware(cors.Config{
	// 	Origins:         "*",
	// 	Methods:         "GET, PUT, POST, DELETE",
	// 	RequestHeaders:  "Origin, Authorization, Content-Type",
	// 	ExposedHeaders:  "",
	// 	MaxAge:          50 * time.Second,
	// 	Credentials:     false,
	// 	ValidateHeaders: false,
	// }))

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

func ClearSessionHandler(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Save()

}
