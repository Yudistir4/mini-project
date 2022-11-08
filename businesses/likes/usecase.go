package likes

type LikeUsecase struct {
	likeRepository Repository
}

func NewLikeUsecase(likeRepository Repository) Usecase {

	return &LikeUsecase{likeRepository: likeRepository}
}

func (u *LikeUsecase) Create(domain *Domain) (Domain, error) {
	return u.likeRepository.Create(domain)
}
func (u *LikeUsecase) GetByID(id string) (Domain, error) {
	post, err := u.likeRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}

func (u *LikeUsecase) GetAll(postID string) ([]Domain, error) {
	return u.likeRepository.GetAll(postID)
}

func (u *LikeUsecase) Delete(userID, postID string) error {
	return u.likeRepository.Delete(userID, postID)
}
func (u *LikeUsecase) DeleteAllLikeByPostID(id string) error {
	return u.likeRepository.DeleteAllLikeByPostID(id)
}
