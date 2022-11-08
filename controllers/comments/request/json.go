package request

import (
	"mini-project/businesses/comments"

	"github.com/go-playground/validator/v10"
)

type Comment struct {
	Comment string `json:"comment" validate:"required"`
	PostID  string `json:"post_id" validate:"required"`
	UserID  string
}

func (req *Comment) ToDomain() *comments.Domain {
	return &comments.Domain{
		Comment: req.Comment,
		PostID:  req.PostID,
		UserID:  req.UserID,
	}
}

func (req *Comment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
