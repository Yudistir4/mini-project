package request

import (
	"mini-project/businesses/lecturers"

	"github.com/go-playground/validator/v10"
)

type Lecturer struct {
	Nidn         string `json:"nidn" validate:"required"`
	RumpunBidang string `json:"rumpun_bidang" validate:"required"`
}

func (req *Lecturer) ToDomain() *lecturers.Domain {
	return &lecturers.Domain{
		Nidn:         req.Nidn,
		RumpunBidang: req.RumpunBidang,
	}
}

func (req *Lecturer) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}
