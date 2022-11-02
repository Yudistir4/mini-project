package posts

type PostUsecase struct {
	postRepository Repository
}

func NewPostUsecase(postRepository Repository) Usecase {

	return &PostUsecase{postRepository: postRepository}
}

func (uu *PostUsecase) Create(domain *Domain) Domain {
	return uu.postRepository.Create(domain)
}
func (uu *PostUsecase) GetByID(id string) Domain {
	post := uu.postRepository.GetByID(id)

	if post.ID == "" {
		return Domain{}
	}

	return post
}

func (uu *PostUsecase) GetAll() []Domain {
	return uu.postRepository.GetAll()
}
func (uu *PostUsecase) Update(id string, domain *Domain) Domain {
	return uu.postRepository.Update(id, domain)
}
func (uu *PostUsecase) Delete(id string) bool {
	return uu.postRepository.Delete(id)
}
func (uu *PostUsecase) DeleteAllPostByUserID(id string) bool {
	return uu.postRepository.DeleteAllPostByUserID(id)
}
