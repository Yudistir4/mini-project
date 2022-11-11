package response

import (
	"mini-project/businesses/comments"
	"mini-project/util"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string         `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	Comment string `json:"comment"`
	PostID  string `json:"post_id"`

	UserID     string `json:"user_id"`
	Name       string `json:"name"`
	ProfilePic string `json:"profile_pic"`
}

func FromDomain(domain comments.Domain) Comment {
	if domain.ProfilePic != "" {
		domain.ProfilePic = util.GetConfig("BASE_URL_IMAGES") + domain.ProfilePic
	}
	return Comment{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		Comment: domain.Comment,
		PostID:  domain.PostID,

		UserID:     domain.UserID,
		Name:       domain.Name,
		ProfilePic: domain.ProfilePic,
	}
}
