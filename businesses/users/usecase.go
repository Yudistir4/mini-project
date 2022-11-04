package users

import (
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

func (uu *UserUsecase) CreateUser(domain *Domain) (Domain, error) {

	if domain.UserType == "student" {
		student, err := uu.studentRepository.Create(&students.Domain{
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
		lecturer, err := uu.lecturerRepository.Create(&lecturers.Domain{Nidn: domain.Nidn, RumpunBidang: domain.RumpunBidang})
		if err != nil {
			return Domain{}, err
		}
		domain.LecturerID = lecturer.ID
	}

	return uu.userRepository.CreateUser(domain)
}
func (uu *UserUsecase) Login(domain *Domain) (string, error) {
	user, err := uu.userRepository.GetByEmail(domain)
	if err != nil {
		return "", err
	}

	token := uu.jwtAuth.GenerateToken(user.ID)
	return token, nil
}

func (uu *UserUsecase) GetByID(id string) (Domain, error) {
	return uu.userRepository.GetByID(id)
}
func (uu *UserUsecase) GetAllUsers(limit int, page int, userType string, name string) ([]Domain, error) {
	return uu.userRepository.GetAllUsers(limit, page, userType, name)
}
func (uu *UserUsecase) Update(id string, domain *Domain) (Domain, error) {
	user, err := uu.userRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}

	if domain.UserType == "student" {
		student, err := uu.studentRepository.GetByID(user.StudentID)
		if err != nil {
			return Domain{}, err
		}
		student.Nim = domain.Nim
		student.Angkatan = domain.Angkatan
		student.Semester = domain.Semester
		student.Status = domain.Status
		if err = uu.studentRepository.Update(student.ID, &student); err != nil {
			return Domain{}, err
		}

	} else if domain.UserType == "lecturer" {
		lecturer, err := uu.lecturerRepository.GetByID(user.LecturerID)
		if err != nil {
			return Domain{}, err
		}

		lecturer.Nidn = domain.Nidn
		lecturer.RumpunBidang = domain.RumpunBidang

		if err = uu.lecturerRepository.Update(user.LecturerID, &lecturer); err != nil {

		}

	}
	if err = uu.userRepository.Update(id, domain); err != nil {
		return Domain{}, err
	}
	return uu.userRepository.GetByID(id)
}
func (uu *UserUsecase) Delete(id string) error {
	user, err := uu.userRepository.GetByID(id)
	if err != nil {
		return err
	}

	if uu.postRepository.DeleteAllPostByUserID(id) == false {
		return err
	}
	if err := uu.userRepository.Delete(id); err != nil {
		return err
	}

	if user.ID != "" {
		if user.UserType == "student" {
			uu.studentRepository.Delete(user.StudentID)
		} else if user.UserType == "lecturer" {
			uu.lecturerRepository.Delete(user.LecturerID)
		}
	}

	return nil
}
func (uu *UserUsecase) UpdateProfilePicture(id, filename string) error {
	return uu.userRepository.UpdateProfilePicture(id, filename)
}
