package response

import (
	"mini-project/businesses/comments"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Comment string `json:"comment"`
	PostID  string `json:"post_id"`

	UserType   string `json:"user_type"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	ProfilePic string `json:"profile_pic"`
	Nim        string `json:"nim"`
	Nidn       string `json:"nidn"`
}

func FromDomain(domain comments.Domain) Comment {
	return Comment{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		Comment:    domain.Comment,
		PostID:     domain.PostID,
		UserType:   domain.UserType,
		UserID:     domain.UserID,
		Name:       domain.Name,
		ProfilePic: domain.ProfilePic,
		Nim:        domain.Nim,
	}
}
