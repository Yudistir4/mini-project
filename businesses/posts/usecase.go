package posts

import (
	"mini-project/businesses/comments"
)

type PostUsecase struct {
	postRepository    Repository
	commentRepository comments.Repository
}

func NewPostUsecase(postRepository Repository, commentRepository comments.Repository) Usecase {

	return &PostUsecase{
		postRepository:    postRepository,
		commentRepository: commentRepository,
	}
}

func (u *PostUsecase) Create(domain *Domain) (Domain, error) {
	return u.postRepository.Create(domain)
}
func (u *PostUsecase) GetByID(id string) (Domain, error) {

	post, err := u.postRepository.GetByID(id)
	if err != nil {
		return Domain{}, err
	}
	// count comment
	commentCount, err := u.commentRepository.GetCommentCount(id)
	if err != nil {
		return Domain{}, err
	}
	post.CommentCount = commentCount

	// TODO: count like

	return post, nil
}

func (u *PostUsecase) GetAll(userID string) ([]Domain, error) {

	posts, err := u.postRepository.GetAll(userID)
	if err != nil {
		return []Domain{}, err
	}
	for i := 0; i < len(posts); i++ {
		// count comment
		commentCount, err := u.commentRepository.GetCommentCount(posts[i].ID)
		if err != nil {
			return []Domain{}, err
		}
		posts[i].CommentCount = commentCount

		// TODO: count like
	}
	return posts, nil
}
func (u *PostUsecase) Update(id string, domain *Domain) (Domain, error) {

	return u.postRepository.Update(id, domain)
}
func (u *PostUsecase) Delete(id string) error {
	_, err := u.postRepository.GetByID(id)
	if err != nil {
		return err
	}
	//delete all commment
	if err := u.commentRepository.DeleteAllCommentByPostID(id); err != nil {
		return err
	}
	//TODO: delete all likes
	//TODO: delete all saved
	return u.postRepository.Delete(id)
}
func (u *PostUsecase) DeleteAllPostByUserID(id string) error {
	posts, err := u.postRepository.GetAll(id)
	if err != nil {
		return err
	}
	//TODO: delete all commment
	for _, post := range posts {
		u.commentRepository.DeleteAllCommentByPostID(post.ID)
	}
	//TODO: delete all likes
	//TODO: delete all saved

	return u.postRepository.DeleteAllPostByUserID(id)

}
