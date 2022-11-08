package saves

import (
	"mini-project/businesses/saves"
	"mini-project/drivers/mysql/posts"
	"mini-project/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Save struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	PostID string     `json:"post_id"`
	Post   posts.Post `json:"post"`

	UserID string     `json:"user_id"`
	User   users.User `json:"user"`
}

func FromDomain(domain *saves.Domain) *Save {
	return &Save{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,

		UserID: domain.UserID,
		PostID: domain.PostID,
	}
}
func (rec *Save) ToDomain() saves.Domain {
	return saves.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,

		PostID: rec.PostID,
		UserID: rec.UserID,
	}
}
