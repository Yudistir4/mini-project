package lecturers

type LecturerUsecase struct {
	userRepository Repository
}

func NewLecturerUsecase(userRepository Repository) Usecase {

	return &LecturerUsecase{userRepository: userRepository}
}

func (uu *LecturerUsecase) Create(domain *Domain) Domain {
	return uu.userRepository.Create(domain)
}
func (uu *LecturerUsecase) GetByID(id string) Domain {
	user := uu.userRepository.GetByID(id)

	if user.ID == "" {
		return Domain{}
	}

	return user
}

func (uu *LecturerUsecase) GetAll() []Domain {
	return uu.userRepository.GetAll()
}
func (uu *LecturerUsecase) Update(id string, domain *Domain) Domain {
	return uu.userRepository.Update(id, domain)
}
func (uu *LecturerUsecase) Delete(id string) bool {
	return uu.userRepository.Delete(id)
}
