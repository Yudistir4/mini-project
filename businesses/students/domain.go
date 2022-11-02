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

type Usecase interface {
	Create(userDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, userDomain *Domain) Domain
	Delete(id string) bool
	GetAll() []Domain
}

type Repository interface {
	Create(userDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, userDomain *Domain) Domain
	Delete(id string) bool
	GetAll() []Domain
}
