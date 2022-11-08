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

	LikeCount    int
	CommentCount int

	IsLiked bool
	IsSaved bool

	UserID     string
	UserType   string
	Name       string
	ProfilePic string
	Nim        string
	Nidn       string
}

type Usecase interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(userIDAccessing, postID string) (Domain, error)
	Update(id string, Domain *Domain) (Domain, error)
	Delete(id string) error
	DeleteAllPostByUserID(userID string) error
	GetAll(userIDAccessing, userID string) ([]Domain, error)
	SavePost(userIDAccessing, postID string) error
	UnsavePost(userIDAccessing, postID string) error
	LikePost(userIDAccessing, postID string) error
	UnlikePost(userIDAccessing, postID string) error
}

type Repository interface {
	Create(Domain *Domain) (Domain, error)
	GetByID(postID string) (Domain, error)
	Update(id string, Domain *Domain) (Domain, error)
	Delete(id string) error
	DeleteAllPostByUserID(userID string) error
	GetAll(userID string) ([]Domain, error)
}
