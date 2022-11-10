package users_test

import (
	"errors"
	"mini-project/app/middlewares"
	"mini-project/businesses/lecturers"
	_lecturerMock "mini-project/businesses/lecturers/mocks"
	_postMock "mini-project/businesses/posts/mocks"
	"mini-project/businesses/students"
	_studentMock "mini-project/businesses/students/mocks"
	"mini-project/businesses/users"
	_userMock "mini-project/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepo     _userMock.Repository
	studentsRepo  _studentMock.Repository
	lecturersRepo _lecturerMock.Repository
	postsRepo     _postMock.Repository
	usersService  users.Usecase

	usersDomain     users.Domain
	userUniversitas users.Domain
	userStudent     users.Domain
	userLecturer    users.Domain
	student         students.Domain
	lecturer        lecturers.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUsecase(&usersRepo, &studentsRepo, &lecturersRepo, &postsRepo, &middlewares.ConfigJwt{})

	usersDomain = users.Domain{
		Email:    "user1@mail.com",
		Password: "123456",
	}

	userUniversitas = users.Domain{
		ID:       "123123",
		Email:    "universitas1@mail.com",
		Password: "123456",
		UserType: "university",
		Name:     "Universitas 1",
	}
	userStudent = users.Domain{
		ID:        "21212121",
		Email:     "student1@mail.com",
		Password:  "123456",
		UserType:  "student",
		Name:      "mahasiswa 1",
		StudentID: student.ID,
	}
	userLecturer = users.Domain{
		ID:         "9191919",
		Email:      "lecturer1@mail.com",
		Password:   "123456",
		UserType:   "lecturer",
		Name:       "dosen 1",
		LecturerID: lecturer.ID,
	}
	student = students.Domain{
		ID:       "101010",
		Nim:      "20190801226",
		Angkatan: 2020,
		Semester: 2019,
	}
	lecturer = lecturers.Domain{
		ID:           "101010",
		Nidn:         "20191818",
		RumpunBidang: "teknik informasi",
	}

	m.Run()
}

func TestCreateUser(t *testing.T) {
	t.Run("CreateUser | Success", func(t *testing.T) {
		usersRepo.On("GetByID", userUniversitas.ID).Return(userUniversitas, nil).Once()
		usersRepo.On("CreateUser", &usersDomain).Return(usersDomain, nil).Once()

		result, _ := usersService.CreateUser(userUniversitas.ID, &usersDomain)
		assert.NotNil(t, result)
	})

	t.Run("CreateUser | Failed (User Not Authenticated) ", func(t *testing.T) {
		usersRepo.On("GetByID", "").Return(users.Domain{}, errors.New("Not Authenticated")).Once()
		usersRepo.On("CreateUser", &usersDomain).Return(usersDomain, nil).Once()

		_, err := usersService.CreateUser("", &users.Domain{})
		assert.Equal(t, err.Error(), "User Not Authenticated")
	})
	t.Run("CreateUser | Failed (Create Student Fail) ", func(t *testing.T) {
		usersRepo.On("GetByID", userUniversitas.ID).Return(userUniversitas, nil).Once()

		studentsRepo.On("Create", &students.Domain{}).Return(students.Domain{}, errors.New("err")).Once()
		_, err := usersService.CreateUser(userUniversitas.ID, &users.Domain{UserType: "student"})
		assert.Equal(t, err.Error(), "err")
	})
	t.Run("CreateUser | Failed (Create Lecturer Fail) ", func(t *testing.T) {
		usersRepo.On("GetByID", userUniversitas.ID).Return(userUniversitas, nil).Once()

		lecturersRepo.On("Create", &lecturers.Domain{}).Return(lecturers.Domain{}, errors.New("err")).Once()
		_, err := usersService.CreateUser(userUniversitas.ID, &users.Domain{UserType: "lecturer"})
		assert.Equal(t, err.Error(), "err")
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepo.On("GetByEmail", &usersDomain).Return(users.Domain{}, nil).Once()

		result, _ := usersService.Login(&usersDomain)
		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepo.On("GetByEmail", &users.Domain{}).Return(users.Domain{}, errors.New("invalid")).Once()

		_, err := usersService.Login(&users.Domain{})

		assert.NotNil(t, err)
	})

}
func TestGetByID(t *testing.T) {
	t.Run("GetByID | Valid", func(t *testing.T) {
		usersRepo.On("GetByID", "").Return(users.Domain{}, nil).Once()
		_, err := usersService.GetByID("")
		assert.Nil(t, err)
	})

	t.Run("GetByID | InValid", func(t *testing.T) {
		usersRepo.On("GetByID", "").Return(users.Domain{}, errors.New("invalid")).Once()
		_, err := usersService.GetByID("")
		assert.NotNil(t, err)
	})

}
func TestGetAllUsers(t *testing.T) {
	t.Run("GetAllUsers | Valid", func(t *testing.T) {
		usersRepo.On("GetAllUsers", 10, 1, "", "").Return([]users.Domain{}, nil).Once()
		_, err := usersService.GetAllUsers(10, 1, "", "")
		assert.Nil(t, err)
	})

	t.Run("GetAllUsers | InValid", func(t *testing.T) {
		usersRepo.On("GetAllUsers", 10, 1, "", "").Return([]users.Domain{}, errors.New("invalid")).Once()
		_, err := usersService.GetAllUsers(10, 1, "", "")
		assert.NotNil(t, err)
	})

}
func TestUpdate(t *testing.T) {
	t.Run("Update UserType Universitas | Valid", func(t *testing.T) {
		usersRepo.On("GetByID", userUniversitas.ID).Return(userUniversitas, nil).Once()
		usersRepo.On("Update", userUniversitas.ID, &userUniversitas).Return(nil).Once()
		usersRepo.On("GetByID", userUniversitas.ID).Return(userUniversitas, nil).Once()
		_, err := usersService.Update(userUniversitas.ID, &userUniversitas)
		assert.Nil(t, err)
	})

}
func TestDelete(t *testing.T) {
	t.Run("Delete UserType Universitas | Valid", func(t *testing.T) {
		usersRepo.On("GetByID", userUniversitas.ID).Return(userUniversitas, nil).Once()
		postsRepo.On("DeleteAllPostByUserID", userUniversitas.ID).Return(nil).Once()
		usersRepo.On("Delete", userUniversitas.ID).Return(nil).Once()

		err := usersService.Delete(userUniversitas.ID)
		assert.Nil(t, err)
	})
	t.Run("Delete UserType Student | Valid", func(t *testing.T) {
		usersRepo.On("GetByID", userStudent.ID).Return(userStudent, nil).Once()
		postsRepo.On("DeleteAllPostByUserID", userStudent.ID).Return(nil).Once()
		usersRepo.On("Delete", userStudent.ID).Return(nil).Once()

		studentsRepo.On("Delete", "").Return(nil).Once()
		err := usersService.Delete(userStudent.ID)
		assert.Nil(t, err)
	})
	t.Run("Delete UserType Lecturer | Valid", func(t *testing.T) {
		usersRepo.On("GetByID", userLecturer.ID).Return(userLecturer, nil).Once()
		postsRepo.On("DeleteAllPostByUserID", userLecturer.ID).Return(nil).Once()
		usersRepo.On("Delete", userLecturer.ID).Return(nil).Once()

		lecturersRepo.On("Delete", "").Return(nil).Once()
		err := usersService.Delete(userLecturer.ID)
		assert.Nil(t, err)
	})
	t.Run("Delete User | Invalid", func(t *testing.T) {
		usersRepo.On("GetByID", userLecturer.ID).Return(userLecturer, errors.New("User Not Found")).Once()

		err := usersService.Delete(userLecturer.ID)
		assert.NotNil(t, err)
	})

}
func TestUpdateProfilePicture(t *testing.T) {
	t.Run("Update Profile Picture | Valid", func(t *testing.T) {
		usersRepo.On("UpdateProfilePicture", userUniversitas.ID, "example.png").Return(nil).Once()

		err := usersService.UpdateProfilePicture(userUniversitas.ID, "example.png")
		assert.Nil(t, err)
	})

}
