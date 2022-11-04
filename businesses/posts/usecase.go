package posts

type PostUsecase struct {
	postRepository Repository
}

func NewPostUsecase(postRepository Repository) Usecase {

	return &PostUsecase{postRepository: postRepository}
}

func (uu *PostUsecase) Create(domain *Domain) (Domain, error) {
	return uu.postRepository.Create(domain)
}
func (uu *PostUsecase) GetByID(id string) (Domain, error) {
	post, err := uu.postRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}

func (uu *PostUsecase) GetAll() ([]Domain, error) {
	return uu.postRepository.GetAll()
}
func (uu *PostUsecase) Update(id string, domain *Domain) (Domain, error) {
	return uu.postRepository.Update(id, domain)
}
func (uu *PostUsecase) Delete(id string) error {
	return uu.postRepository.Delete(id)
}
func (uu *PostUsecase) DeleteAllPostByUserID(id string) error {
	return uu.postRepository.DeleteAllPostByUserID(id)
}
