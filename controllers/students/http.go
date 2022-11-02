package students

import (
	"mini-project/businesses/students"
	"mini-project/controllers"
	"mini-project/controllers/students/request"
	"mini-project/controllers/students/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type StudentController struct {
	studentUsecase students.Usecase
}

func NewStudentController(authUC students.Usecase) *StudentController {
	return &StudentController{studentUsecase: authUC}
}

func (ctrl *StudentController) Create(c echo.Context) error {

	studentInput := request.Student{}

	if err := c.Bind(&studentInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := studentInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	student := ctrl.studentUsecase.Create(studentInput.ToDomain())
	return c.JSON(http.StatusCreated, response.FromDomain(student))
}
func (ctrl *StudentController) GetById(c echo.Context) error {
	id := c.Param("id")

	student := ctrl.studentUsecase.GetByID(id)

	return controllers.NewResponse(c, http.StatusOK, "success", "get student", response.FromDomain(student))

}
func (ctrl *StudentController) GetAll(c echo.Context) error {

	studentsData := ctrl.studentUsecase.GetAll()

	students := []response.Student{}

	for _, student := range studentsData {
		students = append(students, response.FromDomain(student))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all students", students)

}

func (ctrl *StudentController) UpdateStudent(c echo.Context) error {
	id := c.Param("id")
	studentInput := request.Student{}

	if err := c.Bind(&studentInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := studentInput.Validate()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	student := ctrl.studentUsecase.Update(id, studentInput.ToDomain())
	return controllers.NewResponse(c, http.StatusOK, "success", "Update Student", response.FromDomain(student))
}
func (ctrl *StudentController) DeleteStudent(c echo.Context) error {
	id := c.Param("id")

	status := ctrl.studentUsecase.Delete(id)
	if status == false {
		return controllers.NewResponse(c, http.StatusOK, "failed", "Delete Student", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Delete Student", "")
}
