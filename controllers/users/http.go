package users

import (
	"fmt"
	"mini-project/businesses/users"
	"mini-project/controllers"
	"mini-project/controllers/users/request"
	"mini-project/controllers/users/response"
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

	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	user := ctrl.userUsecase.CreateUser(userInput.ToDomain())
	return c.JSON(http.StatusCreated, response.FromDomain(user))
}
func (ctrl *UserController) Login(c echo.Context) error {
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	token := ctrl.userUsecase.Login(userInput.ToDomain())

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})

	}
	return c.JSON(http.StatusCreated, map[string]string{
		"token": token,
	})

}
func (ctrl *UserController) GetAllUsers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))

	fmt.Println("page:", page)
	fmt.Println("limit:", limit)

	usersData := ctrl.userUsecase.GetAllUsers(limit, page)

	users := []response.User{}

	for _, user := range usersData {
		users = append(users, response.FromDomain(user))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all users", users)

}
func (ctrl *UserController) GetByID(c echo.Context) error {

	id := c.Param("id")
	user := ctrl.userUsecase.GetByID(id)
	if user.ID == "" {

		return controllers.NewResponse(c, http.StatusNotFound, "failed", "user Not Found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all users", response.FromDomain(user))

}

func (ctrl *UserController) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	userInput := request.User{}

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := userInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	user := ctrl.userUsecase.Update(id, userInput.ToDomain())

	if user.ID == "" {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", "User Not Found", "")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Update User", response.FromDomain(user))
}
func (ctrl *UserController) DeleteUser(c echo.Context) error {
	id := c.Param("id")

	status := ctrl.userUsecase.Delete(id)
	if status == false {

		return controllers.NewResponse(c, http.StatusNotFound, "failed", "User Not found", status)
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "Delete User", status)
}
