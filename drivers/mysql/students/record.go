package students

import (
	"mini-project/businesses/students"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`
	Nim       string         `json:"nim" gorm:"unique"`
	Angkatan  int            `json:"angkatan"`
	Semester  int            `json:"semester"`
	Status    string         `json:"status"`
}

func FromDomain(domain *students.Domain) *Student {
	return &Student{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Nim:       domain.Nim,
		Angkatan:  domain.Angkatan,
		Semester:  domain.Semester,
		Status:    domain.Status,
	}
}
func (rec *Student) ToDomain() students.Domain {
	return students.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
		Nim:       rec.Nim,
		Angkatan:  rec.Angkatan,
		Semester:  rec.Semester,
		Status:    rec.Status,
	}
}
