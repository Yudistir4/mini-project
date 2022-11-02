package response

import (
	"mini-project/businesses/students"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Nim      string `json:"nim"`
	Angkatan int    `json:"angkatan"`
	Semester int    `json:"semester"`
	Status   string `json:"status"`
}

func FromDomain(domain students.Domain) Student {
	return Student{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		Nim:      domain.Nim,
		Angkatan: domain.Angkatan,
		Semester: domain.Semester,
		Status:   domain.Status,
	}
}
