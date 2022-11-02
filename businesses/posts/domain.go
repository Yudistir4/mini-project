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
	FileType string
	Caption  string

	UserID     string
	UserType   string
	Name       string
	ProfilePic string
	Nim        string
	Nidn       string
}

type Usecase interface {
	Create(Domain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, Domain *Domain) Domain
	Delete(id string) bool
	DeleteAllPostByUserID(userID string) bool
	GetAll() []Domain
}

type Repository interface {
	Create(Domain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, Domain *Domain) Domain
	Delete(id string) bool
	DeleteAllPostByUserID(userID string) bool
	GetAll() []Domain
}
