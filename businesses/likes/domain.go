package likes

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	PostID string

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

	Delete(userID, postID string) error
	DeleteAllLikeByPostID(postID string) error
	// TODO: pagination
	GetAll(postID string) ([]Domain, error)
}

type Repository interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	GetByUserIDAndPostID(userID, postID string) (Domain, error)
	Delete(userID, postID string) error
	DeleteAllLikeByPostID(postID string) error
	GetAll(postID string) ([]Domain, error)
	GetLikeCount(postID string) (int, error)
}
