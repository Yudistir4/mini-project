package posts

import (
	"mini-project/businesses/posts"
	"mini-project/controllers"
	"mini-project/controllers/posts/request"
	"mini-project/controllers/posts/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostController struct {
	lecturerUsecase posts.Usecase
}

func NewPostController(authUC posts.Usecase) *PostController {
	return &PostController{lecturerUsecase: authUC}
}

func (ctrl *PostController) Create(c echo.Context) error {

	dataInput := request.Post{}

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
func (ctrl *PostController) GetById(c echo.Context) error {
	id := c.Param("id")

	data := ctrl.lecturerUsecase.GetByID(id)

	return controllers.NewResponse(c, http.StatusOK, "success", "get data", response.FromDomain(data))

}
func (ctrl *PostController) GetAll(c echo.Context) error {

	postsData := ctrl.lecturerUsecase.GetAll()

	posts := []response.Post{}

	for _, data := range postsData {
		posts = append(posts, response.FromDomain(data))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all posts", posts)

}

func (ctrl *PostController) UpdatePost(c echo.Context) error {
	id := c.Param("id")
	dataInput := request.Post{}

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
	return controllers.NewResponse(c, http.StatusOK, "success", "Update Post", response.FromDomain(data))
}
func (ctrl *PostController) DeletePost(c echo.Context) error {
	id := c.Param("id")

	status := ctrl.lecturerUsecase.Delete(id)
	if status == false {
		return controllers.NewResponse(c, http.StatusOK, "failed", "Delete Post", "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Delete Post", "")
}
