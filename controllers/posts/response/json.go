package response

import (
	"mini-project/businesses/posts"
	"mini-project/util"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        string         `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

	CommentCount int    `json:"comment_count"`
	LikeCount    int    `json:"like_count"`
	FileName     string `json:"file_name"`
	Caption      string `json:"caption"`
	UserType     string `json:"user_type"`
	UserID       string `json:"user_id"`
	Name         string `json:"name"`
	ProfilePic   string `json:"profile_pic"`
	Nim          string `json:"nim"`
	Nidn         string `json:"nidn"`

	IsSaved bool `json:"is_saved"`
	IsLiked bool `json:"is_liked"`
}

func FromDomain(domain posts.Domain) Post {
	if domain.ProfilePic != "" {
		domain.ProfilePic = util.GetConfig("BASE_URL_IMAGES") + domain.ProfilePic
	}
	if domain.FileName != "" {
		domain.ProfilePic = util.GetConfig("BASE_URL_IMAGES") + domain.FileName
	}
	return Post{
		ID:           domain.ID,
		CreatedAt:    domain.CreatedAt,
		UpdatedAt:    domain.UpdatedAt,
		DeletedAt:    domain.DeletedAt,
		CommentCount: domain.CommentCount,
		LikeCount:    domain.LikeCount,

		FileName:   domain.FileName,
		Caption:    domain.Caption,
		UserType:   domain.UserType,
		UserID:     domain.UserID,
		Name:       domain.Name,
		ProfilePic: domain.ProfilePic,
		Nim:        domain.Nim,

		IsSaved: domain.IsSaved,
		IsLiked: domain.IsLiked,
	}
}
