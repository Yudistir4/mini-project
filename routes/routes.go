package routes

import (
	"mini-project/controllers/comments"
	"mini-project/controllers/posts"
	"mini-project/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware     middleware.JWTConfig
	UserController    users.UserController
	PostController    posts.PostController
	CommentController comments.CommentController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	auth := e.Group("api/v1/auth")

	auth.POST("/signup", cl.UserController.CreateUser)
	auth.POST("/login", cl.UserController.Login)

	users := e.Group("api/v1/users", middleware.JWTWithConfig(cl.JWTMiddleware))
	users.GET("", cl.UserController.GetAllUsers)
	users.GET("/:id", cl.UserController.GetByID)
	users.DELETE("/:id", cl.UserController.DeleteUser)
	users.PUT("/:id", cl.UserController.UpdateUser)
	users.PUT("/:id/profilepicture", cl.UserController.UpdateProfilePicture)

	posts := e.Group("api/v1/posts", middleware.JWTWithConfig(cl.JWTMiddleware))
	posts.GET("", cl.PostController.GetAll)
	posts.GET("/:id", cl.PostController.GetById)
	posts.POST("", cl.PostController.Create)
	posts.DELETE("/:id", cl.PostController.DeletePost)
	posts.PUT("/:id", cl.PostController.UpdatePost)

	comments := e.Group("api/v1/comments", middleware.JWTWithConfig(cl.JWTMiddleware))
	comments.GET("", cl.CommentController.GetAll)
	comments.GET("/:id", cl.CommentController.GetById)
	comments.POST("", cl.CommentController.Create)
	comments.DELETE("/:id", cl.CommentController.DeletePost)

}
