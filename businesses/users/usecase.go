package users

import (
	"mini-project/app/middlewares"
)

type UserUsecase struct {
	userRepository Repository
	// studentRepository students.Repository
	jwtAuth *middlewares.ConfigJwt
}

func NewUserUsecase(userRepository Repository, jwtAuth *middlewares.ConfigJwt) Usecase {

	return &UserUsecase{userRepository: userRepository, jwtAuth: jwtAuth}
}

func (uu *UserUsecase) CreateUser(domain *Domain) Domain {

	return uu.userRepository.CreateUser(domain)
}
func (uu *UserUsecase) Login(domain *Domain) string {
	user := uu.userRepository.GetByEmail(domain)

	if user.ID == "" {
		return ""
	}

	token := uu.jwtAuth.GenerateToken(user.ID)
	return token
}

func (uu *UserUsecase) GetByID(id string) Domain {
	return uu.userRepository.GetByID(id)
}
func (uu *UserUsecase) GetAllUsers(limit, page int) []Domain {
	return uu.userRepository.GetAllUsers(limit, page)
}
func (uu *UserUsecase) Update(id string, domain *Domain) Domain {
	return uu.userRepository.Update(id, domain)
}
func (uu *UserUsecase) Delete(id string) bool {
	return uu.userRepository.Delete(id)
}
