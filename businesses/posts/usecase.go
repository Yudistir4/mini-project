package posts

import (
	"mini-project/businesses/comments"
	"mini-project/businesses/likes"
	"mini-project/businesses/saves"
)

type PostUsecase struct {
	postRepository    Repository
	commentRepository comments.Repository
	likeRepository    likes.Repository
	saveRepository    saves.Repository
}

func NewPostUsecase(postRepository Repository, commentRepository comments.Repository, likeRepository likes.Repository, saveRepository saves.Repository) Usecase {

	return &PostUsecase{
		postRepository:    postRepository,
		commentRepository: commentRepository,
		likeRepository:    likeRepository,
		saveRepository:    saveRepository,
	}
}

func (u *PostUsecase) Create(domain *Domain) (Domain, error) {
	return u.postRepository.Create(domain)
}
func (u *PostUsecase) GetByID(userIDAccessing, postID string) (Domain, error) {

	post, err := u.postRepository.GetByID(postID)
	if err != nil {
		return Domain{}, err
	}
	// count comment
	commentCount, err := u.commentRepository.GetCommentCount(postID)
	if err != nil {
		return Domain{}, err
	}
	post.CommentCount = commentCount

	// count like
	likeCount, err := u.likeRepository.GetLikeCount(postID)
	if err != nil {
		return Domain{}, err
	}
	post.LikeCount = likeCount

	// Get is saved
	if err := u.saveRepository.GetByUserIDAndPostID(userIDAccessing, postID); err == nil {
		post.IsSaved = true
	}

	//  Get is Liked
	if err := u.likeRepository.GetByUserIDAndPostID(userIDAccessing, postID); err == nil {
		post.IsLiked = true
	}

	return post, nil
}

func (u *PostUsecase) GetAll(userIDAccessing, userID string) ([]Domain, error) {

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

		// count like
		likeCount, err := u.likeRepository.GetLikeCount(posts[i].ID)
		if err != nil {
			return []Domain{}, err
		}
		posts[i].LikeCount = likeCount

		// is saved
		if err := u.saveRepository.GetByUserIDAndPostID(userIDAccessing, posts[i].ID); err == nil {
			posts[i].IsSaved = true
		}

		// todo: is liked
		if err := u.likeRepository.GetByUserIDAndPostID(userIDAccessing, posts[i].ID); err == nil {
			posts[i].IsLiked = true
		}
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
	// delete all likes
	if err := u.likeRepository.DeleteAllLikeByPostID(id); err != nil {
		return err
	}

	// delete all saved
	if err := u.saveRepository.DeleteAllSaveByPostID(id); err != nil {
		return err
	}
	return u.postRepository.Delete(id)
}
func (u *PostUsecase) DeleteAllPostByUserID(id string) error {
	posts, err := u.postRepository.GetAll(id)
	if err != nil {
		return err
	}
	for _, post := range posts {
		// delete all commment
		u.commentRepository.DeleteAllCommentByPostID(post.ID)
		// delete all likes
		u.likeRepository.DeleteAllLikeByPostID(post.ID)
		// delete all saved
		u.saveRepository.DeleteAllSaveByPostID(post.ID)

	}

	return u.postRepository.DeleteAllPostByUserID(id)

}

func (u *PostUsecase) SavePost(userIDAccessing, postID string) error {

	return u.saveRepository.Create(userIDAccessing, postID)

}
func (u *PostUsecase) UnsavePost(userIDAccessing, postID string) error {

	return u.saveRepository.Delete(userIDAccessing, postID)

}
func (u *PostUsecase) LikePost(userIDAccessing, postID string) error {
	return u.likeRepository.Create(userIDAccessing, postID)
}
func (u *PostUsecase) UnlikePost(userIDAccessing, postID string) error {

	return u.likeRepository.Delete(userIDAccessing, postID)

}
