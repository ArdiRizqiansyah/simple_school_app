package route

import (
	"be-classroom/app/config"
	"be-classroom/app/database"
	"be-classroom/handler"
	"be-classroom/repository/classroom_repository/classroom_pg"
	"be-classroom/repository/classroom_student_repository/classroom_student_pg"
	"be-classroom/repository/student_repository/student_pg"
	"be-classroom/repository/user_repository/user_pg"
	"be-classroom/service/auth_service"
	"be-classroom/service/classroom_service"
	"be-classroom/service/classroom_student_service"
	"be-classroom/service/student_service"
	"be-classroom/service/user_service"
	"fmt"

	"github.com/gin-gonic/gin"
)

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		// Log to ensure middleware is running
		fmt.Println("CORS middleware executed")

		c.Next()
	}
}

func StartApp() {
	config.LoadEnv()

	database.InitializeDatabase()
	db := database.GetInstanceDatabaseConnection()

	userRepo := user_pg.NewUserRepository(db)
	userService := user_service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	classroomRepo := classroom_pg.NewClassroomRepository(db)
	classroomService := classroom_service.NewClassroomService(classroomRepo)
	classroomHandler := handler.NewClassroomHandler(classroomService)

	studentRepo := student_pg.NewStudentRepository(db)
	studentService := student_service.NewStudentService(studentRepo)
	studentHandler := handler.NewStudentHandler(studentService)

	classroomStudentRepo := classroom_student_pg.NewClassroomStudentRepository(db)
	classroomStudentService := classroom_student_service.NewClassroomStudentService(classroomStudentRepo, classroomRepo)
	classroomStudentHandler := handler.NewClassroomStudentHandler(classroomStudentService, studentService)

	dashboardHandler := handler.NewDashboardHandler(classroomService, studentService)

	authService := auth_service.NewAuthService(userRepo)

	app := gin.Default()

	// Apply CORS middleware to allow specific origin
	app.Use(CORSMiddleware())

	// routing
	users := app.Group("users")
	{
		users.POST("/register", userHandler.Register)
		users.POST("/login", userHandler.Login)
		users.GET("/profile", authService.Authentication(), userHandler.Profile)

		// dashboard
		users.GET("/dashboard", authService.Authentication(), dashboardHandler.Dashboard)
	}

	classrooms := app.Group("classrooms")
	{
		classrooms.Use(authService.Authentication())

		// classrooms
		classrooms.GET("", classroomHandler.GetAllClassrooms)
		classrooms.GET("/:classroomId", classroomHandler.GetClassroomById)
		classrooms.POST("", classroomHandler.CreateClassroom)
		classrooms.PUT("/:classroomId", classroomHandler.EditClassroom)
		classrooms.DELETE("/:classroomId", classroomHandler.DeleteClassroom)

		// classrooms/students
		classrooms.GET("/:classroomId/students", classroomStudentHandler.GetAllClassroomStudents)
		classrooms.POST("/:classroomId/students", classroomStudentHandler.AddStudentToClassroom)
		classrooms.DELETE("/:classroomId/students/:classroomStudentId", classroomStudentHandler.DeleteClassroomStudent)
	}

	students := app.Group("students")
	{
		students.Use(authService.Authentication())

		students.GET("", studentHandler.GetAllStudents)
		students.GET("/:studentId", studentHandler.GetStudentById)
		students.POST("", studentHandler.CreateStudent)
		students.PUT("/:studentId", studentHandler.EditStudent)
		students.DELETE("/:studentId", studentHandler.DeleteStudent)
	}

	app.Run(":" + config.AppConfig().Port)
}
