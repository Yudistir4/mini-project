package lecturers

import (
	"mini-project/businesses/lecturers"
	"mini-project/controllers"
	"mini-project/controllers/lecturers/request"
	"mini-project/controllers/lecturers/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LecturerController struct {
	lecturerUsecase lecturers.Usecase
}

func NewLecturerController(authUC lecturers.Usecase) *LecturerController {
	return &LecturerController{lecturerUsecase: authUC}
}

func (ctrl *LecturerController) Create(c echo.Context) error {

	dataInput := request.Lecturer{}

	if err := c.Bind(&dataInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := dataInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	data := ctrl.lecturerUsecase.Create(dataInput.ToDomain())
	return controllers.NewResponse(c, http.StatusCreated, "success", "create lecturer", response.FromDomain(data))
}
func (ctrl *LecturerController) GetById(c echo.Context) error {
	id := c.Param("id")

	data := ctrl.lecturerUsecase.GetByID(id)

	return controllers.NewResponse(c, http.StatusOK, "success", "get data", response.FromDomain(data))

}
func (ctrl *LecturerController) GetAll(c echo.Context) error {

	lecturersData := ctrl.lecturerUsecase.GetAll()

	lecturers := []response.Lecturer{}

	for _, data := range lecturersData {
		lecturers = append(lecturers, response.FromDomain(data))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all lecturers", lecturers)

}

func (ctrl *LecturerController) UpdateLecturer(c echo.Context) error {
	id := c.Param("id")
	dataInput := request.Lecturer{}

	if err := c.Bind(&dataInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := dataInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	data := ctrl.lecturerUsecase.Update(id, dataInput.ToDomain())
	return controllers.NewResponse(c, http.StatusOK, "success", "Update Lecturer", response.FromDomain(data))
}
func (ctrl *LecturerController) DeleteLecturer(c echo.Context) error {
	id := c.Param("id")

	status := ctrl.lecturerUsecase.Delete(id)
	if status == false {
		return controllers.NewResponse(c, http.StatusOK, "failed", "Delete Lecturer", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Delete Lecturer", "")
}
