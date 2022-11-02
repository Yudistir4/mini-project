package students

type StudentUsecase struct {
	userRepository Repository
}

func NewStudentUsecase(userRepository Repository) Usecase {

	return &StudentUsecase{userRepository: userRepository}
}

func (uu *StudentUsecase) Create(domain *Domain) Domain {
	return uu.userRepository.Create(domain)
}
func (uu *StudentUsecase) GetByID(id string) Domain {
	user := uu.userRepository.GetByID(id)

	if user.ID == "" {
		return Domain{}
	}

	return user
}

func (uu *StudentUsecase) GetAll() []Domain {
	return uu.userRepository.GetAll()
}
func (uu *StudentUsecase) Update(id string, domain *Domain) Domain {
	return uu.userRepository.Update(id, domain)
}
func (uu *StudentUsecase) Delete(id string) bool {
	return uu.userRepository.Delete(id)
}
