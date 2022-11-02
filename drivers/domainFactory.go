package drivers

import (
	"mini-project/businesses/lecturers"
	"mini-project/businesses/posts"
	"mini-project/businesses/students"
	"mini-project/businesses/users"

	lecturersDB "mini-project/drivers/mysql/lecturers"
	postsDB "mini-project/drivers/mysql/posts"
	studentsDB "mini-project/drivers/mysql/students"
	usersDB "mini-project/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) users.Repository {
	return usersDB.NewMySQLRepository(conn)
}
func NewStudentRepository(conn *gorm.DB) students.Repository {
	return studentsDB.NewMySQLRepository(conn)
}
func NewLecturerRepository(conn *gorm.DB) lecturers.Repository {
	return lecturersDB.NewMySQLRepository(conn)
}

func NewPostRepository(conn *gorm.DB) posts.Repository {
	return postsDB.NewMySQLRepository(conn)
}
