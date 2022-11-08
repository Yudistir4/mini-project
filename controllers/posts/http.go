package posts

import (
	"mini-project/businesses/posts"
	"mini-project/controllers"
	"mini-project/controllers/posts/request"
	"mini-project/controllers/posts/response"
	"mini-project/util"
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
	filename, err := util.FileHandling(c)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")
	}

	dataInput := request.Post{}
	dataInput.FileName = filename
	if err := c.Bind(&dataInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	if err := dataInput.Validate(); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})

	}

	data, err := ctrl.lecturerUsecase.Create(dataInput.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")

	}
	return controllers.NewResponse(c, http.StatusCreated, "success", "create post", response.FromDomain(data))
}
func (ctrl *PostController) GetById(c echo.Context) error {
	id := c.Param("id")

	data, err := ctrl.lecturerUsecase.GetByID(id)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")

	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get data", response.FromDomain(data))

}
func (ctrl *PostController) GetAll(c echo.Context) error {
	userID := c.QueryParam("user_id")
	postsData, err := ctrl.lecturerUsecase.GetAll(userID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")

	}

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

	data, err := ctrl.lecturerUsecase.Update(id, dataInput.ToDomain())
	if err != nil {

		return controllers.NewResponse(c, http.StatusOK, "failed", err.Error(), "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Update Post", response.FromDomain(data))
}
func (ctrl *PostController) DeletePost(c echo.Context) error {
	id := c.Param("id")

	err := ctrl.lecturerUsecase.Delete(id)
	if err != nil {
		return controllers.NewResponse(c, http.StatusOK, "failed", err.Error(), "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Delete Post", "")
}
