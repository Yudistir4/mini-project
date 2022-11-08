package main

import (
	"log"
	"mini-project/app/middlewares"
	"mini-project/routes"

	_commentUseCase "mini-project/businesses/comments"
	_commentController "mini-project/controllers/comments"

	_postUseCase "mini-project/businesses/posts"
	_postController "mini-project/controllers/posts"

	_userUseCase "mini-project/businesses/users"
	_userController "mini-project/controllers/users"

	"mini-project/drivers"
	"mini-project/drivers/mysql"
	"mini-project/util"

	"github.com/labstack/echo/v4"
)

func main() {
	configDB := mysql.ConfigDB{
		DB_USERNAME: util.GetConfig("DB_USERNAME"),
		DB_PASSWORD: util.GetConfig("DB_PASSWORD"),
		DB_HOST:     util.GetConfig("DB_HOST"),
		DB_PORT:     util.GetConfig("DB_PORT"),
		DB_NAME:     util.GetConfig("DB_NAME"),
	}

	db := configDB.InitDB()

	mysql.DBMigrate(db)

	configJWT := middlewares.ConfigJwt{
		SecretJWT:       util.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 240,
	}

	e := echo.New()

	saveRepo := drivers.NewSaveRepository(db)

	likeRepo := drivers.NewLikeRepository(db)

	commentRepo := drivers.NewCommentRepository(db)
	commentUsecase := _commentUseCase.NewCommentUsecase(commentRepo)
	commentController := _commentController.NewCommentController(commentUsecase)

	postRepo := drivers.NewPostRepository(db)
	postUsecase := _postUseCase.NewPostUsecase(postRepo, commentRepo, likeRepo, saveRepo)
	postController := _postController.NewPostController(postUsecase)

	lecturerRepo := drivers.NewLecturerRepository(db)
	studentRepo := drivers.NewStudentRepository(db)

	userRepo := drivers.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, studentRepo, lecturerRepo, postRepo, &configJWT)
	userController := _userController.NewUserController(userUsecase)

	routesInit := routes.ControllerList{
		JWTMiddleware:     configJWT.Init(),
		UserController:    *userController,
		PostController:    *postController,
		CommentController: *commentController,
		// LikeController:    *likeController,
		// SaveController:    *saveController,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}
