package users

import (
	"fmt"
	"mini-project/app/middlewares"
	"mini-project/businesses/users"
	"mini-project/controllers"
	"mini-project/controllers/users/request"
	"mini-project/controllers/users/response"
	"mini-project/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(authUC users.Usecase) *UserController {
	return &UserController{userUsecase: authUC}
}

func (ctrl *UserController) CreateUser(c echo.Context) error {
	userIDAccessing := middlewares.GetUserIDFromToken(c)

	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Invalid Request", "")

	}

	if err := userInput.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Validation failed", "")
	}

	user, err := ctrl.userUsecase.CreateUser(userIDAccessing, userInput.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}
	return controllers.NewResponse(c, http.StatusCreated, "success", "create user", response.FromDomain(user))
}
func (ctrl *UserController) Login(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Invalid Request", "")

	}

	if err := userInput.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Validation failed", "")
	}

	token, err := ctrl.userUsecase.Login(userInput.ToDomain())

	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")

	}
	return controllers.NewResponse(c, http.StatusOK, "success", "login success", token)

}
func (ctrl *UserController) GetAllUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	userType := c.QueryParam("userType")
	name := c.QueryParam("name")

	fmt.Println("page:", page)
	fmt.Println("limit:", limit)

	usersData, err := ctrl.userUsecase.GetAllUsers(limit, page, userType, name)
	if err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", err.Error(), "")
	}

	users := []response.User{}

	for _, user := range usersData {
		users = append(users, response.FromDomain(user))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all users", users)

}
func (ctrl *UserController) GetByID(c echo.Context) error {

	id := c.Param("id")
	user, err := ctrl.userUsecase.GetByID(id)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get users", response.FromDomain(user))

}

func (ctrl *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")

	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Invalid Request", "")
	}

	if err := userInput.Validate(); err != nil {
		return controllers.NewResponse(c, http.StatusBadRequest, "failed", "Validation failed", "")
	}

	user, err := ctrl.userUsecase.Update(id, userInput.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Update User", response.FromDomain(user))
}
func (ctrl *UserController) UpdateProfilePicture(c echo.Context) error {
	id := middlewares.GetUserIDFromToken(c)
	filename, err := util.FileHandling(c)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}
	if err := ctrl.userUsecase.UpdateProfilePicture(id, filename); err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Update Profile Picture", "")
}
func (ctrl *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	if err := ctrl.userUsecase.Delete(id); err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Delete User", "")
}
