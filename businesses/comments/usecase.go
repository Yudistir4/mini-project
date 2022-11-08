package comments

type CommentUsecase struct {
	commentRepository Repository
}

func NewCommentUsecase(commentRepository Repository) Usecase {

	return &CommentUsecase{commentRepository: commentRepository}
}

func (u *CommentUsecase) Create(domain *Domain) (Domain, error) {
	return u.commentRepository.Create(domain)
}
func (u *CommentUsecase) GetByID(id string) (Domain, error) {
	post, err := u.commentRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}

func (u *CommentUsecase) GetAll(postID string) ([]Domain, error) {
	return u.commentRepository.GetAll(postID)
}

func (u *CommentUsecase) Delete(id string) error {
	return u.commentRepository.Delete(id)
}
func (u *CommentUsecase) DeleteAllCommentByPostID(id string) error {
	return u.commentRepository.DeleteAllCommentByPostID(id)
}
