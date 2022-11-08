package comments

type CommentUsecase struct {
	commentRepository Repository
}

func NewCommentUsecase(commentRepository Repository) Usecase {

	return &CommentUsecase{commentRepository: commentRepository}
}

func (uu *CommentUsecase) Create(domain *Domain) (Domain, error) {
	return uu.commentRepository.Create(domain)
}
func (uu *CommentUsecase) GetByID(id string) (Domain, error) {
	post, err := uu.commentRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}

func (uu *CommentUsecase) GetAll(postID string) ([]Domain, error) {
	return uu.commentRepository.GetAll(postID)
}

func (uu *CommentUsecase) Delete(id string) error {
	return uu.commentRepository.Delete(id)
}
func (uu *CommentUsecase) DeleteAllCommentByPostID(id string) error {
	return uu.commentRepository.DeleteAllCommentByPostID(id)
}
