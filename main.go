package main

import (
	"log"
	"mini-project/app/middlewares"
	"mini-project/routes"

	_postUseCase "mini-project/businesses/posts"
	_postController "mini-project/controllers/posts"

	_userUseCase "mini-project/businesses/users"
	_userController "mini-project/controllers/users"

	_lecturerUseCase "mini-project/businesses/lecturers"
	_lecturerController "mini-project/controllers/lecturers"

	_studentUseCase "mini-project/businesses/students"
	_studentController "mini-project/controllers/students"

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

	postRepo := drivers.NewPostRepository(db)
	postUsecase := _postUseCase.NewPostUsecase(postRepo)
	postController := _postController.NewPostController(postUsecase)

	lecturerRepo := drivers.NewLecturerRepository(db)
	lecturerUsecase := _lecturerUseCase.NewLecturerUsecase(lecturerRepo)
	lecturerController := _lecturerController.NewLecturerController(lecturerUsecase)

	studentRepo := drivers.NewStudentRepository(db)
	studentUsecase := _studentUseCase.NewStudentUsecase(studentRepo)
	studentController := _studentController.NewStudentController(studentUsecase)

	userRepo := drivers.NewUserRepository(db)
	userUsecase := _userUseCase.NewUserUsecase(userRepo, postRepo, &configJWT)
	userController := _userController.NewUserController(userUsecase)

	routesInit := routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserController:     *userController,
		StudentController:  *studentController,
		LecturerController: *lecturerController,
		PostController:     *postController,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(":8000"))
}
