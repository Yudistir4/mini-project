package posts

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	FileName string
	Caption  string

	UserID     string
	UserType   string
	Name       string
	ProfilePic string
	Nim        string
	Nidn       string
}

type Usecase interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, Domain *Domain) (Domain, error)
	Delete(id string) error
	DeleteAllPostByUserID(userID string) error
	GetAll() ([]Domain, error)
}

type Repository interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, Domain *Domain) (Domain, error)
	Delete(id string) error
	DeleteAllPostByUserID(userID string) error
	GetAll() ([]Domain, error)
}
