package users

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	Email     string
	Password  string

	UserType      string
	Name          string
	ProfilePic    string
	Bio           string
	EksternalLink string
	Instagram     string
	Linkedin      string
	Whatsapp      string

	StudentID string
	Nim       string
	Angkatan  int
	Semester  int
	Status    string

	LecturerID   string
	Nidn         string
	RumpunBidang string
}

type Usecase interface {
	CreateUser(userDomain *Domain) Domain
	Login(userDomain *Domain) string
	GetByID(id string) Domain
	Update(id string, userDomain *Domain) Domain
	Delete(id string) bool
	GetAllUsers(limit, page int) []Domain
}

type Repository interface {
	CreateUser(userDomain *Domain) Domain
	GetByEmail(userDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, userDomain *Domain) Domain
	Delete(id string) bool
	GetAllUsers(limit, page int) []Domain
}
