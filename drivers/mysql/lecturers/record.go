package lecturers

import (
	"mini-project/businesses/lecturers"
	"time"

	"gorm.io/gorm"
)

type Lecturer struct {
	ID           string         `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `json:"deletedAt"`
	RumpunBidang string         `json:"rumpun_bidang"`
	Nidn         string         `json:"nidn"`
}

func FromDomain(domain *lecturers.Domain) *Lecturer {
	return &Lecturer{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		RumpunBidang: domain.RumpunBidang,
		Nidn:         domain.Nidn,
	}
}
func (rec *Lecturer) ToDomain() lecturers.Domain {
	return lecturers.Domain{
		ID:           rec.ID,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
		RumpunBidang: rec.RumpunBidang,
		Nidn:         rec.Nidn,
	}
}
