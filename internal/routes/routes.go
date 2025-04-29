package routes

import (
	"university_system/internal/auth"
	infraRepo "university_system/internal/infrastructure/repository"
	controller "university_system/internal/university/controllers"
	"university_system/internal/university/services"
	"university_system/pkg/databases"
	"university_system/pkg/middleware"

	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterUserRoutes(router *gin.Engine) {
	userRepo := infraRepo.NewUserRepository(databases.Instance)
	userService := services.NewUserService(userRepo)
	studentRepo := infraRepo.NewStudentRepository(databases.Instance)
	courseRepo := infraRepo.NewCourseRepository(databases.Instance)
	teacherRepo := infraRepo.NewTeacherRepository(databases.Instance)
	managerRepo := infraRepo.NewManagerRepository(databases.Instance)
	markRepo := infraRepo.NewGradeRepository(databases.Instance)
	studentController := controller.NewStudentController(studentRepo)
	userController := controller.NewUserController(userService)
	courseController := controller.NewCourseController(courseRepo)
	teacherController := controller.NewTeacherController(teacherRepo)
	managerController := controller.NewManagerController(managerRepo)
	markController := controller.NewCourseMarkController(markRepo)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/users", middleware.RoleMiddleware("admin", "manager"), userController.GetUsers)
		protected.POST("/users", middleware.RoleMiddleware("admin", "manager"), userController.CreateUser)
		protected.GET("/users/:id", middleware.RoleMiddleware("admin", "manager"), userController.GetUserById)
		protected.PUT("/users/:id", middleware.RoleMiddleware("admin", "manager"), userController.UpdateUser)
		protected.DELETE("/users/:id", middleware.RoleMiddleware("admin"), userController.DeleteUser)
	}

	studentRoutes := router.Group("/students")
	studentRoutes.Use(middleware.AuthMiddleware())
	{
		studentRoutes.GET("/", middleware.RoleMiddleware("admin", "manager", "teacher"), studentController.GetStudents)
		studentRoutes.POST("", middleware.RoleMiddleware("admin", "manager"), studentController.CreateStudent)
		studentRoutes.GET("/:id", middleware.RoleMiddleware("admin", "manager", "teacher", "student"), studentController.GetStudentById)
		studentRoutes.PUT("/:id", middleware.RoleMiddleware("admin", "manager", "student"), studentController.UpdateStudent)
		studentRoutes.DELETE("/:id", middleware.RoleMiddleware("admin", "manager"), studentController.DeleteStudent)
		studentRoutes.POST("/:student_id/courses/:course_id", middleware.RoleMiddleware("admin", "manager", "student"), studentController.EnrollStudentToCourse)
		studentRoutes.GET("/:id/courses", middleware.RoleMiddleware("admin", "manager", "teacher", "student"), studentController.GetStudentCourses)
	}

	courseRoutes := router.Group("/courses")
	courseRoutes.Use(middleware.AuthMiddleware())
	{
		courseRoutes.POST("/", middleware.RoleMiddleware("admin", "manager", "teacher"), courseController.CreateCourse)
		courseRoutes.GET("/:id", middleware.RoleMiddleware("admin", "manager", "teacher", "student"), courseController.GetCourseByID)
		courseRoutes.GET("/", middleware.RoleMiddleware("admin", "manager", "teacher", "student"), courseController.GetAllCourses)
		courseRoutes.PUT("/:id", middleware.RoleMiddleware("admin", "manager", "teacher"), courseController.UpdateCourse)
		courseRoutes.DELETE("/:id", middleware.RoleMiddleware("admin", "manager"), courseController.DeleteCourse)
		courseRoutes.GET("/:id/students", middleware.RoleMiddleware("admin", "manager", "teacher"), courseController.GetCourseStudents)
		courseRoutes.GET("/:id/teachers", middleware.RoleMiddleware("admin", "manager", "teacher"), courseController.GetCourseTeachers)
	}

	teacherRoutes := router.Group("/teachers")
	teacherRoutes.Use(middleware.AuthMiddleware())
	{
		teacherRoutes.GET("/", middleware.RoleMiddleware("admin", "manager", "teacher"), teacherController.GetTeachers)
		teacherRoutes.GET("/:id", middleware.RoleMiddleware("admin", "manager", "teacher"), teacherController.GetTeacherByID)
		teacherRoutes.PUT("/:id", middleware.RoleMiddleware("admin", "manager", "teacher"), teacherController.UpdateTeacher)
		teacherRoutes.DELETE("/:id", middleware.RoleMiddleware("admin", "manager"), teacherController.DeleteTeacher)
		teacherRoutes.POST("/", middleware.RoleMiddleware("admin", "manager"), teacherController.CreateTeacher)
		teacherRoutes.GET("/:id/courses", middleware.RoleMiddleware("admin", "manager", "teacher"), teacherController.GetTeacherCourses)
		teacherRoutes.POST("/:id/courses/:course_id/students/:student_id/PutFirstAtt", middleware.RoleMiddleware("admin", "teacher"), markController.AddFirstAttestation)
		teacherRoutes.POST("/:id/courses/:course_id/students/:student_id/PutSecondAtt", middleware.RoleMiddleware("admin", "teacher"), markController.AddSecondAttestation)
		teacherRoutes.POST("/:id/courses/:course_id/students/:student_id/PutFinalMark", middleware.RoleMiddleware("admin", "teacher"), markController.AddFinalExamMark)

	}

	managerRoutes := router.Group("/managers")
	managerRoutes.Use(middleware.AuthMiddleware())
	{
		managerRoutes.GET("/", middleware.RoleMiddleware("admin", "manager"), managerController.GetManagers)
		managerRoutes.GET("/:id", middleware.RoleMiddleware("admin", "manager"), managerController.GetManagerById)
		managerRoutes.PUT("/:id", middleware.RoleMiddleware("admin", "manager"), managerController.UpdateManager)
		managerRoutes.DELETE("/:id", middleware.RoleMiddleware("admin"), managerController.DeleteManager)
		managerRoutes.POST("/", middleware.RoleMiddleware("admin", "manager"), managerController.CreateManager)
		managerRoutes.POST("/:id/teachers/:teacher_id/courses/:course_id", middleware.RoleMiddleware("admin", "manager"), managerController.AssignTeacherToCourse)
	}

	marksRoutes := router.Group("/marks")
	marksRoutes.Use(middleware.AuthMiddleware())
	{
		marksRoutes.GET("/student/:student_id", middleware.RoleMiddleware("admin", "manager", "teacher", "student"), markController.GetStudentMarks)
		marksRoutes.GET("/course/:course_id", middleware.RoleMiddleware("admin", "manager", "teacher", "student"), markController.GetCourseMarks)
	}

	router.POST("/login", auth.Login)
	router.POST("/refresh", auth.Refresh)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(files.Handler))
}
