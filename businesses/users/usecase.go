package users

import (
	"errors"
	"mini-project/app/middlewares"
	"mini-project/businesses/lecturers"
	"mini-project/businesses/posts"
	"mini-project/businesses/students"
)

type UserUsecase struct {
	userRepository     Repository
	studentRepository  students.Repository
	lecturerRepository lecturers.Repository
	postRepository     posts.Repository
	jwtAuth            *middlewares.ConfigJwt
}

func NewUserUsecase(userRepository Repository, studentRepository students.Repository, lecturerRepository lecturers.Repository, postRepository posts.Repository, jwtAuth *middlewares.ConfigJwt) Usecase {

	return &UserUsecase{userRepository: userRepository, studentRepository: studentRepository, lecturerRepository: lecturerRepository, postRepository: postRepository, jwtAuth: jwtAuth}
}

func (u *UserUsecase) CreateUser(userIDAccessing string, domain *Domain) (Domain, error) {
	user, _ := u.userRepository.GetByID(userIDAccessing)
	if user.UserType != "university" {
		return Domain{}, errors.New("User Not Authenticated")
	}
	if domain.UserType == "student" {
		student, err := u.studentRepository.Create(&students.Domain{
			Nim:      domain.Nim,
			Angkatan: domain.Angkatan,
			Semester: domain.Semester,
			Status:   domain.Status,
		})
		if err != nil {
			return Domain{}, err
		}
		domain.StudentID = student.ID

	} else if domain.UserType == "lecturer" {
		lecturer, err := u.lecturerRepository.Create(&lecturers.Domain{Nidn: domain.Nidn, RumpunBidang: domain.RumpunBidang})
		if err != nil {
			return Domain{}, err
		}
		domain.LecturerID = lecturer.ID
	}

	return u.userRepository.CreateUser(domain)
}
func (u *UserUsecase) Login(domain *Domain) (string, error) {
	user, err := u.userRepository.GetByEmail(domain)
	if err != nil {
		return "", err
	}

	token := u.jwtAuth.GenerateToken(user.ID)

	return token, nil
}

func (u *UserUsecase) GetByID(id string) (Domain, error) {
	return u.userRepository.GetByID(id)
}
func (u *UserUsecase) GetAllUsers(limit int, page int, userType string, name string) ([]Domain, error) {
	return u.userRepository.GetAllUsers(limit, page, userType, name)
}
func (u *UserUsecase) Update(id string, domain *Domain) (Domain, error) {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}

	if domain.UserType == "student" {
		student, err := u.studentRepository.GetByID(user.StudentID)
		if err != nil {
			return Domain{}, err
		}
		student.Nim = domain.Nim
		student.Angkatan = domain.Angkatan
		student.Semester = domain.Semester
		student.Status = domain.Status
		if err = u.studentRepository.Update(student.ID, &student); err != nil {
			return Domain{}, err
		}

	} else if domain.UserType == "lecturer" {
		lecturer, err := u.lecturerRepository.GetByID(user.LecturerID)
		if err != nil {
			return Domain{}, err
		}

		lecturer.Nidn = domain.Nidn
		lecturer.RumpunBidang = domain.RumpunBidang

		if err = u.lecturerRepository.Update(user.LecturerID, &lecturer); err != nil {

		}

	}
	if err = u.userRepository.Update(id, domain); err != nil {
		return Domain{}, err
	}
	return u.userRepository.GetByID(id)
}
func (u *UserUsecase) Delete(id string) error {
	user, err := u.userRepository.GetByID(id)
	if err != nil {
		return err
	}

	if err := u.postRepository.DeleteAllPostByUserID(id); err != nil {
		return err
	}
	if err := u.userRepository.Delete(id); err != nil {
		return err
	}

	if user.ID != "" {
		if user.UserType == "student" {
			u.studentRepository.Delete(user.StudentID)
		} else if user.UserType == "lecturer" {
			u.lecturerRepository.Delete(user.LecturerID)
		}
	}

	return nil
}
func (u *UserUsecase) UpdateProfilePicture(id, filename string) error {
	return u.userRepository.UpdateProfilePicture(id, filename)
}
