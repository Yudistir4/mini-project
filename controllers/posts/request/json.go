package request

import (
	"mini-project/businesses/posts"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	FileName string `json:"file_name" validate:"required"`
	Caption  string `form:"caption" json:"caption"`
	UserID   string
}

func (req *Post) ToDomain() *posts.Domain {
	return &posts.Domain{
		FileName: req.FileName,
		Caption:  req.Caption,
		UserID:   req.UserID,
	}
}

func (req *Post) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
