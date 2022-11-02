package request

import (
	"mini-project/businesses/posts"

	"github.com/go-playground/validator/v10"
)

type Post struct {
	FileName string `json:"file_name"`
	FileType string `json:"file_type"`
	Caption  string `json:"caption" validate:"required"`
	UserID   string `json:"user_id" validate:"required"`
}

func (req *Post) ToDomain() *posts.Domain {
	return &posts.Domain{
		FileName: req.FileName,
		FileType: req.FileType,
		Caption:  req.Caption,
		UserID:   req.UserID,
	}
}

func (req *Post) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
