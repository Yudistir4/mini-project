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
	UserID string
}

type Repository interface {
	Create(userID, postID string) error
	GetByID(id string) (Domain, error)
	GetByUserIDAndPostID(userID, postID string) error
	Delete(userID, postID string) error
	DeleteAllLikeByPostID(postID string) error
	GetAll(postID string) ([]Domain, error)
	GetLikeCount(postID string) (int, error)
}
