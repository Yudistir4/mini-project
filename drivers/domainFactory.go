package drivers

import (
	"mini-project/businesses/blogs"
	"mini-project/businesses/categories"
	"mini-project/businesses/lecturers"
	"mini-project/businesses/students"
	"mini-project/businesses/users"
	blogsDB "mini-project/drivers/mysql/blogs"
	categoriesDB "mini-project/drivers/mysql/categories"
	lecturersDB "mini-project/drivers/mysql/lecturers"
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
func NewBlogRepository(conn *gorm.DB) blogs.Repository {
	return blogsDB.NewMySQLRepository(conn)
}
func NewCategoryRepository(conn *gorm.DB) categories.Repository {
	return categoriesDB.NewMySQLRepository(conn)
}
