package students

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Nim      string
	Angkatan int
	Semester int
	Status   string
}

type Repository interface {
	Create(userDomain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, userDomain *Domain) error
	Delete(id string) error
}
