package drivers

import (
	"mini-project/businesses/comments"
	"mini-project/businesses/lecturers"
	"mini-project/businesses/likes"
	"mini-project/businesses/posts"
	"mini-project/businesses/saves"
	"mini-project/businesses/students"
	"mini-project/businesses/users"

	commentsDB "mini-project/drivers/mysql/comments"
	lecturersDB "mini-project/drivers/mysql/lecturers"
	likesDB "mini-project/drivers/mysql/likes"
	postsDB "mini-project/drivers/mysql/posts"
	savesDB "mini-project/drivers/mysql/saves"
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
func NewCommentRepository(conn *gorm.DB) comments.Repository {
	return commentsDB.NewMySQLRepository(conn)
}
func NewLikeRepository(conn *gorm.DB) likes.Repository {
	return likesDB.NewMySQLRepository(conn)
}
func NewSaveRepository(conn *gorm.DB) saves.Repository {
	return savesDB.NewMySQLRepository(conn)
}
