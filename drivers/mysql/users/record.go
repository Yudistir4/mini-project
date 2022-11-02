package users

import (
	"mini-project/businesses/users"
	"mini-project/drivers/mysql/lecturers"
	"mini-project/drivers/mysql/students"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            string         `json:"id" gorm:"type:varchar(255);primary_key;"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
	Name          string         `json:"name"`
	ProfilePic    string         `json:"profile_pic"`
	Email         string         `json:"email" gorm:"unique" faker:"email"`
	UserType      string         `json:"user_type"`
	Password      string         `json:"password" faker:"password"`
	Bio           string         `json:"bio"`
	EksternalLink string         `json:"eksternal_link"`
	Instagram     string         `json:"instagram"`
	Linkedin      string         `json:"linkedin"`
	Whatsapp      string         `json:"whatsapp"`

	StudentID string           `json:"student_id"`
	Student   students.Student `json:"student"`

	LecturerID string             `json:"lecturer_id"`
	Lecturer   lecturers.Lecturer `json:"lecturer"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
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
		StudentID:     domain.StudentID,
		LecturerID:    domain.LecturerID,

		Student: students.Student{
			Nim:      domain.Nim,
			Angkatan: domain.Angkatan,
			Semester: domain.Semester,
			Status:   domain.Status,
		},
		Lecturer: lecturers.Lecturer{
			Nidn:         domain.Nidn,
			RumpunBidang: domain.RumpunBidang,
		},
	}
}
func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:            rec.ID,
		CreatedAt:     rec.CreatedAt,
		UpdatedAt:     rec.UpdatedAt,
		DeletedAt:     rec.DeletedAt,
		Email:         rec.Email,
		Password:      rec.Password,
		UserType:      rec.UserType,
		Name:          rec.Name,
		ProfilePic:    rec.ProfilePic,
		Bio:           rec.Bio,
		EksternalLink: rec.EksternalLink,
		Instagram:     rec.Instagram,
		Linkedin:      rec.Linkedin,
		Whatsapp:      rec.Whatsapp,

		StudentID: rec.StudentID,
		Nim:       rec.Student.Nim,
		Angkatan:  rec.Student.Angkatan,
		Semester:  rec.Student.Semester,
		Status:    rec.Student.Status,

		LecturerID:   rec.LecturerID,
		Nidn:         rec.Lecturer.Nidn,
		RumpunBidang: rec.Lecturer.RumpunBidang,
	}
}
