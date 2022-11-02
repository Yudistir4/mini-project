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

type Usecase interface {
	Create(Domain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, Domain *Domain) Domain
	Delete(id string) bool
	GetAll() []Domain
}

type Repository interface {
	Create(Domain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, Domain *Domain) Domain
	Delete(id string) bool
	GetAll() []Domain
}
