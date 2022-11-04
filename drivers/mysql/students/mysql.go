package students

import (
	"mini-project/businesses/students"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type StudentRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) students.Repository {
	return &StudentRepository{
		conn: conn,
	}
}

func (ur *StudentRepository) Create(domain *students.Domain) (students.Domain, error) {
	student := FromDomain(domain)
	student.ID = uuid.New().String()
	if err := ur.conn.Save(&student).Error; err != nil {
		return students.Domain{}, err
	}

	return student.ToDomain(), nil
}
func (ur *StudentRepository) GetByID(id string) (students.Domain, error) {
	var student Student
	if err := ur.conn.First(&student, "id = ?", id).Error; err != nil {

		return students.Domain{}, err
	}

	return student.ToDomain(), nil
}

func (ur *StudentRepository) Update(id string, domain *students.Domain) error {

	updateStudent := FromDomain(domain)

	if err := ur.conn.Save(&updateStudent).Error; err != nil {
		return err
	}
	return nil
}
func (ur *StudentRepository) Delete(id string) error {
	var student Student
	if err := ur.conn.Delete(&student, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
