package routes

import (
	"university_system/internal/auth"
	controller "university_system/internal/university/controllers"
	"university_system/internal/university/repository"
	"university_system/pkg/databases"
	"university_system/pkg/middleware"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRepo := repository.NewUserRepository(databases.Instance)
	studentRepo := repository.NewStudentRepository(databases.Instance)
	courseRepo := repository.NewCourseRepository(databases.Instance)
	teacherRepo := repository.NewTeacherRepository(databases.Instance)
	managerRepo := repository.NewManagerRepository(databases.Instance)
	markRepo := repository.NewCourseMarkRepository(databases.Instance)
	studentController := controller.NewStudentController(studentRepo)
	userController := controller.NewUserController(userRepo)
	courseController := controller.NewCourseController(courseRepo)
	teacherController := controller.NewTeacherController(teacherRepo)
	managerController := controller.NewManagerController(managerRepo)
	markController := controller.NewCourseMarkController(markRepo)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", userController.GetUsers)
		protected.GET("/users/:id", userController.GetUserById)
		protected.PUT("/users/:id", userController.UpdateUser)
		protected.DELETE("/users/:id", userController.DeleteUser)
	}

	studentRoutes := router.Group("/students")
	studentRoutes.Use(middleware.AuthMiddleware())
	{
		studentRoutes.GET("/", studentController.GetStudents)
		studentRoutes.POST("", studentController.CreateStudent)
		studentRoutes.GET("/:id", studentController.GetStudentById)
		studentRoutes.PUT("/:id", studentController.UpdateStudent)
		studentRoutes.DELETE("/:id", studentController.DeleteStudent)
		studentRoutes.POST("/:student_id/courses/:course_id", studentController.EnrollStudentToCourse)
		studentRoutes.GET("/:id/courses", studentController.GetStudentCourses)
	}

	courseRoutes := router.Group("/courses")
	courseRoutes.Use(middleware.AuthMiddleware())
	{
		courseRoutes.POST("/", courseController.CreateCourse)
		courseRoutes.GET("/:id", courseController.GetCourseByID)
		courseRoutes.GET("/", courseController.GetAllCourses)
		courseRoutes.PUT("/:id", courseController.UpdateCourse)
		courseRoutes.DELETE("/:id", courseController.DeleteCourse)
		courseRoutes.GET("/:id/students", courseController.GetCourseStudents)
		courseRoutes.GET("/:id/teachers", courseController.GetCourseTeachers)
	}

	teacherRoutes := router.Group("/teachers")
	teacherRoutes.Use(middleware.AuthMiddleware())
	{
		teacherRoutes.GET("/", teacherController.GetTeachers)
		teacherRoutes.GET("/:id", teacherController.GetTeacherByID)
		teacherRoutes.PUT("/:id", teacherController.UpdateTeacher)
		teacherRoutes.DELETE("/:id", teacherController.DeleteTeacher)
		teacherRoutes.POST("/", teacherController.CreateTeacher)
		teacherRoutes.GET("/:id/courses", teacherController.GetTeacherCourses)
		teacherRoutes.POST("/:id/courses/:course_id/students/:student_id/PutFirstAtt", markController.AddFirstAttestation)
		teacherRoutes.POST("/:id/courses/:course_id/students/:student_id/PusSecondAtt", markController.AddSecondAttestation)
		teacherRoutes.POST("/:id/courses/:course_id/students/:student_id/PutFinalMark", markController.AddFirstAttestation)

	}

	managerRoutes := router.Group("/managers")
	managerRoutes.Use(middleware.AuthMiddleware())
	{
		managerRoutes.GET("/", managerController.GetManagers)
		managerRoutes.GET("/:id", managerController.GetManagerById)
		managerRoutes.PUT("/:id", managerController.UpdateManager)
		managerRoutes.DELETE("/:id", managerController.DeleteManager)
		managerRoutes.POST("/", managerController.CreateManager)
		managerRoutes.POST("/:id/teachers/:teacher_id/courses/:course_id", managerController.AssignTeacherToCourse)
	}

	marksRoutes := router.Group("/marks")
	managerRoutes.Use(middleware.AuthMiddleware())
	{
		marksRoutes.GET("/student/:student_id", markController.GetStudentMarks)
		marksRoutes.GET("/course/:course_id", markController.GetCourseMarks)
	}

	router.POST("/register", userController.Register)
	router.POST("/login", auth.Login)
	router.POST("/refresh", auth.Refresh)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}
