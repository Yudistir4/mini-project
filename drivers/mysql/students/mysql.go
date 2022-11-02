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

func (ur *StudentRepository) Create(domain *students.Domain) students.Domain {
	student := FromDomain(domain)
	student.ID = uuid.New().String()
	ur.conn.Save(&student)

	return student.ToDomain()
}
func (ur *StudentRepository) GetByID(id string) students.Domain {
	var student Student
	ur.conn.First(&student, "id = ?", id)

	if student.ID == "" {
		return students.Domain{}
	}

	return student.ToDomain()
}
func (ur *StudentRepository) GetAll() []students.Domain {
	var rec []Student
	ur.conn.Find(&rec)

	studentsDomain := []students.Domain{}

	for _, student := range rec {
		studentsDomain = append(studentsDomain, student.ToDomain())
	}

	return studentsDomain
}
func (ur *StudentRepository) Update(id string, domain *students.Domain) students.Domain {
	student := ur.GetByID(id)

	updateStudent := FromDomain(&student)
	updateStudent.Nim = domain.Nim
	updateStudent.Angkatan = domain.Angkatan
	updateStudent.Semester = domain.Semester
	updateStudent.Status = domain.Status

	if err := ur.conn.Save(&updateStudent).Error; err != nil {
		return students.Domain{}
	}
	return updateStudent.ToDomain()
}
func (ur *StudentRepository) Delete(id string) bool {
	var student Student
	if err := ur.conn.Delete(&student, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}
