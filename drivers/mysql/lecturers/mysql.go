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

func (ur *LecturerRepository) Create(domain *lecturers.Domain) (lecturers.Domain, error) {
	lecturer := FromDomain(domain)
	lecturer.ID = uuid.New().String()
	if err := ur.conn.Save(&lecturer).Error; err != nil {
		return lecturers.Domain{}, err
	}

	return lecturer.ToDomain(), nil
}
func (ur *LecturerRepository) GetByID(id string) (lecturers.Domain, error) {
	var lecturer Lecturer
	if err := ur.conn.First(&lecturer, "id = ?", id).Error; err != nil {
		return lecturers.Domain{}, err
	}
	return lecturer.ToDomain(), nil
}

func (ur *LecturerRepository) Update(id string, domain *lecturers.Domain) error {

	updateLecturer := FromDomain(domain)

	if err := ur.conn.Save(&updateLecturer).Error; err != nil {
		return err
	}
	return nil
}
func (ur *LecturerRepository) Delete(id string) error {
	var lecturer Lecturer
	if err := ur.conn.Delete(&lecturer, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
