package posts

import (
	"mini-project/businesses/posts"
	"mini-project/drivers/mysql/users"
	"time"

	"gorm.io/gorm"
)

type Post struct {
	ID        string         `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt"`

	FileName string `json:"file_name"`
	Caption  string `json:"caption"`

	UserID string     `json:"user_id"`
	User   users.User `json:"user"`
}

func FromDomain(domain *posts.Domain) *Post {
	return &Post{
		ID:        domain.ID,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
		FileName:  domain.FileName,
		Caption:   domain.Caption,
		UserID:    domain.UserID,
	}
}
func (rec *Post) ToDomain() posts.Domain {
	return posts.Domain{
		ID:        rec.ID,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,

		FileName: rec.FileName,
		Caption:  rec.Caption,
		UserID:   rec.UserID,

		UserType:   rec.User.UserType,
		Name:       rec.User.Name,
		ProfilePic: rec.User.ProfilePic,
		Nim:        rec.User.Student.Nim,
		Nidn:       rec.User.Lecturer.Nidn,
	}
}
