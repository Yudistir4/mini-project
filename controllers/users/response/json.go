package response

import (
	"mini-project/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string         `json:"id" gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Email         string         `json:"email"`
	Password      string         `json:"password"`
	UserType      string         `json:"user_type"`
	Name          string         `json:"name"`
	ProfilePic    string         `json:"profile_pic"`
	Bio           string         `json:"bio"`
	EksternalLink string         `json:"eksternal_link"`
	Instagram     string         `json:"instagram"`
	Linkedin      string         `json:"linkedin"`
	Whatsapp      string         `json:"whatsapp"`

	StudentID string `json:"student_id"`
	Nim       string `json:"nim"`
	Angkatan  int    `json:"angkatan"`
	Semester  int    `json:"semester"`
	Status    string `json:"status"`

	LecturerID   string `json:"lecturer_id"`
	Nidn         string `json:"nidn"`
	RumpunBidang string `json:"rumpun_bidang"`
}

func FromDomain(domain users.Domain) User {
	return User{
		ID:            domain.ID,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
		DeletedAt:     domain.DeletedAt,
		Email:         domain.Email,
		Password:      domain.Password,
		UserType:      domain.UserType,
		Name:          domain.Name,
		ProfilePic:    domain.ProfilePic,
		Bio:           domain.Bio,
		EksternalLink: domain.EksternalLink,
		Instagram:     domain.Instagram,
		Linkedin:      domain.Linkedin,
		Whatsapp:      domain.Whatsapp,

		StudentID: domain.StudentID,
		Nim:       domain.Nim,
		Angkatan:  domain.Angkatan,
		Semester:  domain.Semester,
		Status:    domain.Status,

		LecturerID:   domain.LecturerID,
		Nidn:         domain.Nidn,
		RumpunBidang: domain.RumpunBidang,
	}
}
