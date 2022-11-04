package lecturers

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Nidn         string
	RumpunBidang string
}

type Repository interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, Domain *Domain) error
	Delete(id string) error
}
