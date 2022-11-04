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
	CreateUser(userDomain *Domain) (Domain, error)
	Login(userDomain *Domain) (string, error)
	GetByID(id string) (Domain, error)
	Update(id string, userDomain *Domain) (Domain, error)
	Delete(id string) error
	GetAllUsers(limit int, page int, userType string, name string) ([]Domain, error)
	UpdateProfilePicture(id, filename string) error
}

type Repository interface {
	CreateUser(userDomain *Domain) (Domain, error)
	GetByEmail(userDomain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, userDomain *Domain) error
	Delete(id string) error
	GetAllUsers(limit int, page int, userType string, name string) ([]Domain, error)
	UpdateProfilePicture(id, filename string) error
}
