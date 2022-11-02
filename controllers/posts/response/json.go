package response

import (
	"mini-project/businesses/posts"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	FileName   string `json:"file_name"`
	FileType   string `json:"file_type"`
	Caption    string `json:"caption"`
	UserType   string `json:"user_type"`
	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	ProfilePic string `json:"profile_pic"`
	Nim        string `json:"nim"`
	Nidn       string `json:"nidn"`
}

func FromDomain(domain posts.Domain) Post {
	return Post{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		FileName:   domain.FileName,
		FileType:   domain.FileType,
		Caption:    domain.Caption,
		UserType:   domain.UserType,
		UserID:     domain.UserID,
		Name:       domain.Name,
		ProfilePic: domain.ProfilePic,
		Nim:        domain.Nim,
	}
}
