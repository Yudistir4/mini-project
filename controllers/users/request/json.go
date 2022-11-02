package request

import (
	"mini-project/businesses/users"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Email         string `json:"email" validate:"required,email"`
	Password      string `json:"password" validate:"required"`
	UserType      string `json:"user_type"`
	Name          string `json:"name"`
	ProfilePic    string `json:"profile_pic"`
	Bio           string `json:"bio"`
	EksternalLink string `json:"eksternal_link"`
	Instagram     string `json:"instagram"`
	Linkedin      string `json:"linkedin"`
	Whatsapp      string `json:"whatsapp"`

	StudentID  string `json:"student_id"`
	LecturerID string `json:"lecturer_id"`
}

func (req *User) ToDomain() *users.Domain {
	return &users.Domain{
		Email:         req.Email,
		Password:      req.Password,
		UserType:      req.UserType,
		Name:          req.Name,
		ProfilePic:    req.ProfilePic,
		Bio:           req.Bio,
		EksternalLink: req.EksternalLink,
		Instagram:     req.Instagram,
		Linkedin:      req.Linkedin,
		Whatsapp:      req.Whatsapp,
		StudentID:     req.StudentID,
		LecturerID:    req.LecturerID,
	}
}

func (req *User) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
