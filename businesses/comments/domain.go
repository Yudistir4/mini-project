package comments

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt

	Comment    string
	PostID     string
	UserID     string
	Name       string
	ProfilePic string
}

type Usecase interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Delete(id string) error
	DeleteAllCommentByPostID(postID string) error
	GetAll(postID string) ([]Domain, error)
}

type Repository interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Delete(id string) error
	DeleteAllCommentByPostID(postID string) error
	GetAll(postID string) ([]Domain, error)
	GetCommentCount(postID string) (int, error)
}
