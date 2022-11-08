package likes

import (
	"mini-project/businesses/likes"
	"mini-project/controllers"
	"mini-project/controllers/likes/request"
	"mini-project/controllers/likes/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LikeController struct {
	likeUsecase likes.Usecase
}

func NewLikeController(authUC likes.Usecase) *LikeController {
	return &LikeController{likeUsecase: authUC}
}

func (ctrl *LikeController) Create(c echo.Context) error {

	dataInput := request.Like{}

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

	data, err := ctrl.likeUsecase.Create(dataInput.ToDomain())
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")

	}
	return controllers.NewResponse(c, http.StatusCreated, "success", "create like", response.FromDomain(data))
}
func (ctrl *LikeController) GetById(c echo.Context) error {
	id := c.Param("id")

	data, err := ctrl.likeUsecase.GetByID(id)
	if err != nil {
		return controllers.NewResponse(c, http.StatusNotFound, "failed", err.Error(), "")

	}

	return controllers.NewResponse(c, http.StatusOK, "success", "get data", response.FromDomain(data))

}
func (ctrl *LikeController) GetAll(c echo.Context) error {
	postID := c.QueryParam("postID")
	likesData, err := ctrl.likeUsecase.GetAll(postID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusInternalServerError, "failed", err.Error(), "")

	}

	likes := []response.Like{}

	for _, data := range likesData {
		likes = append(likes, response.FromDomain(data))
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "all likes", likes)

}

func (ctrl *LikeController) DeleteLike(c echo.Context) error {
	dataInput := request.Like{}

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

	err := ctrl.likeUsecase.Delete(dataInput.UserID, dataInput.PostID)
	if err != nil {
		return controllers.NewResponse(c, http.StatusOK, "failed", err.Error(), "")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "Delete Like", "")
}
