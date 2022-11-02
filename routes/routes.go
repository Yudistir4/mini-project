package routes

import (
	"mini-project/controllers/lecturers"
	"mini-project/controllers/posts"
	"mini-project/controllers/students"
	"mini-project/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserController     users.UserController
	StudentController  students.StudentController
	LecturerController lecturers.LecturerController
	PostController     posts.PostController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("api/v1/users")

	users.POST("/signup", cl.UserController.CreateUser)
	users.POST("/login", cl.UserController.Login)

	students := e.Group("api/v1/students")
	students.POST("", cl.StudentController.Create)
	students.GET("", cl.StudentController.GetAll)
	students.GET("/:id", cl.StudentController.GetById)
	students.PUT("/:id", cl.StudentController.UpdateStudent)
	students.DELETE("/:id", cl.StudentController.DeleteStudent)

	lecturers := e.Group("api/v1/lecturers")
	lecturers.POST("", cl.LecturerController.Create)
	lecturers.GET("", cl.LecturerController.GetAll)
	lecturers.GET("/:id", cl.LecturerController.GetById)
	lecturers.PUT("/:id", cl.LecturerController.UpdateLecturer)
	lecturers.DELETE("/:id", cl.LecturerController.DeleteLecturer)

	auth := e.Group("api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	auth.GET("", cl.UserController.GetAllUsers)
	auth.GET("/:id", cl.UserController.GetByID)
	auth.DELETE("/:id", cl.UserController.DeleteUser)
	auth.PUT("/:id", cl.UserController.UpdateUser)

	posts := e.Group("api/v1/posts", middleware.JWTWithConfig(cl.JWTMiddleware))
	posts.GET("", cl.PostController.GetAll)
	posts.GET("/:id", cl.PostController.GetById)
	posts.POST("", cl.PostController.Create)
	posts.DELETE("/:id", cl.PostController.DeletePost)
	posts.PUT("/:id", cl.PostController.UpdatePost)
}
