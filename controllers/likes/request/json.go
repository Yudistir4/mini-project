package request

import (
	"mini-project/businesses/likes"

	"github.com/go-playground/validator/v10"
)

type Like struct {
	PostID string `json:"post_id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
}

func (req *Like) ToDomain() *likes.Domain {
	return &likes.Domain{

		PostID: req.PostID,
		UserID: req.UserID,
	}
}

func (req *Like) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
