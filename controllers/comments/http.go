package comments

import (
	"mini-project/businesses/comments"
	"mini-project/controllers"
	"mini-project/controllers/comments/request"
	"mini-project/controllers/comments/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CommentController struct {
	lecturerUsecase comments.Usecase
}

func NewCommentController(authUC comments.Usecase) *CommentController {
	return &CommentController{lecturerUsecase: authUC}
}

func (ctrl *CommentController) Create(c echo.Context) error {

	dataInput := request.Comment{}

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
func (ctrl *CommentController) GetById(c echo.Context) error {
	id := c.Param("id")

	data, err := ctrl.lecturerUsecase.GetByID(id)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")

	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get data", response.FromDomain(data))

}
func (ctrl *CommentController) GetAll(c echo.Context) error {
	postID := c.QueryParam("postID")
	commentsData, err := ctrl.lecturerUsecase.GetAll(postID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")

	}

	comments := []response.Comment{}

	for _, data := range commentsData {
		comments = append(comments, response.FromDomain(data))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all comments", comments)

}

func (ctrl *CommentController) DeleteComment(c echo.Context) error {
	id := c.Param("id")

	err := ctrl.lecturerUsecase.Delete(id)
	if err != nil {
		return controllers.NewResponse(c, http.StatusOK, "failed", err.Error(), "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Delete Comment", "")
}
