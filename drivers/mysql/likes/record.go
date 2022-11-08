package likes

import (
	"mini-project/businesses/likes"
	"mini-project/drivers/mysql/posts"
	"mini-project/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Like struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	PostID string     `json:"post_id"`
	Post   posts.Post `json:"post"`

	UserID string     `json:"user_id"`
	User   users.User `json:"user"`
}

func FromDomain(domain *likes.Domain) *Like {
	return &Like{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		UserID: domain.UserID,
		PostID: domain.PostID,
	}
}
func (rec *Like) ToDomain() likes.Domain {
	return likes.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,

		PostID: rec.PostID,
		UserID: rec.UserID,
	}
}
