package response

import (
	"mini-project/businesses/lecturers"
	"time"

	"gorm.io/gorm"
)

type Lecturer struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Nidn         string `json:"nidn"`
	RumpunBidang string `json:"rumpun_bidang"`
}

func FromDomain(domain lecturers.Domain) Lecturer {
	return Lecturer{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		Nidn:         domain.Nidn,
		RumpunBidang: domain.RumpunBidang,
	}
}
