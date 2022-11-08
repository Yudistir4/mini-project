package response

import (
	"mini-project/businesses/likes"
	"time"

	"gorm.io/gorm"
)

type Like struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	PostID string `json:"post_id"`

	UserType   string `json:"user_type"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	ProfilePic string `json:"profile_pic"`
	Nim        string `json:"nim"`
	Nidn       string `json:"nidn"`
}

func FromDomain(domain likes.Domain) Like {
	return Like{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		PostID:     domain.PostID,
		UserType:   domain.UserType,
		UserID:     domain.UserID,
		Name:       domain.Name,
		ProfilePic: domain.ProfilePic,
		Nim:        domain.Nim,
	}
}
