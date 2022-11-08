package comments

import (
	"mini-project/businesses/comments"
	"mini-project/drivers/mysql/posts"
	"mini-project/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	Comment string `json:"comment"`

	PostID string     `json:"post_id"`
	Post   posts.Post `json:"post"`

	UserID string     `json:"user_id"`
	User   users.User `json:"user"`
}

func FromDomain(domain *comments.Domain) *Comment {
	return &Comment{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		Comment:   domain.Comment,
		UserID:    domain.UserID,
		PostID:    domain.PostID,
	}
}
func (rec *Comment) ToDomain() comments.Domain {
	return comments.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,

		Comment: rec.Comment,

		PostID: rec.PostID,

		UserID:     rec.UserID,
		UserType:   rec.User.UserType,
		Name:       rec.User.Name,
		ProfilePic: rec.User.ProfilePic,
		Nim:        rec.User.Student.Nim,
		Nidn:       rec.User.Lecturer.Nidn,
	}
}
