package lecturers

import (
	"mini-project/businesses/lecturers"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LecturerRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) lecturers.Repository {
	return &LecturerRepository{
		conn: conn,
	}
}

func (ur *LecturerRepository) Create(domain *lecturers.Domain) lecturers.Domain {
	lecturer := FromDomain(domain)
	lecturer.ID = uuid.New().String()
	ur.conn.Save(&lecturer)

	return lecturer.ToDomain()
}
func (ur *LecturerRepository) GetByID(id string) lecturers.Domain {
	var lecturer Lecturer
	ur.conn.First(&lecturer, "id = ?", id)

	if lecturer.ID == "" {
		return lecturers.Domain{}
	}

	return lecturer.ToDomain()
}
func (ur *LecturerRepository) GetAll() []lecturers.Domain {
	var rec []Lecturer
	ur.conn.Find(&rec)

	lecturersDomain := []lecturers.Domain{}

	for _, lecturer := range rec {
		lecturersDomain = append(lecturersDomain, lecturer.ToDomain())
	}

	return lecturersDomain
}
func (ur *LecturerRepository) Update(id string, domain *lecturers.Domain) lecturers.Domain {
	lecturer := ur.GetByID(id)

	updateLecturer := FromDomain(&lecturer)
	updateLecturer.RumpunBidang = domain.RumpunBidang
	updateLecturer.Nidn = domain.Nidn

	if err := ur.conn.Save(&updateLecturer).Error; err != nil {
		return lecturers.Domain{}
	}
	return updateLecturer.ToDomain()
}
func (ur *LecturerRepository) Delete(id string) bool {
	var lecturer Lecturer
	if err := ur.conn.Delete(&lecturer, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}
