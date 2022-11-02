package request

import (
	"mini-project/businesses/students"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	Nim      string `json:"nim" validate:"required"`
	Angkatan int    `json:"angkatan" validate:"required"`
	Semester int    `json:"semester" validate:"required"`
	Status   string `json:"status" validate:"required"`
}

func (req *Student) ToDomain() *students.Domain {
	return &students.Domain{
		Nim:      req.Nim,
		Angkatan: req.Angkatan,
		Semester: req.Semester,
		Status:   req.Status,
	}
}

func (req *Student) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
