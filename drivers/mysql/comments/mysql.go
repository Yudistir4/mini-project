package comments

import (
	"mini-project/businesses/comments"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) comments.Repository {
	return &CommentRepository{
		conn: conn,
	}
}

func (ur *CommentRepository) Create(domain *comments.Domain) (comments.Domain, error) {
	comment := FromDomain(domain)
	comment.ID = uuid.New().String()

	if err := ur.conn.Save(&comment).Error; err != nil {
		return comments.Domain{}, err
	}

	var dataComment Comment
	if err := ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&dataComment, "id = ?", comment.ID).Error; err != nil {
		return comments.Domain{}, err
	}

	return dataComment.ToDomain(), nil
}
func (ur *CommentRepository) GetByID(id string) (comments.Domain, error) {
	var comment Comment
	if err := ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&comment, "id = ?", id).Error; err != nil {
		return comments.Domain{}, err
	}

	return comment.ToDomain(), nil
}
func (ur *CommentRepository) GetAll(postID string) ([]comments.Domain, error) {
	var rec []Comment
	ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").Find(&rec, "post_id = ?", postID)

	commentsDomain := []comments.Domain{}

	for _, comment := range rec {
		commentsDomain = append(commentsDomain, comment.ToDomain())
	}

	return commentsDomain, nil
}
func (ur *CommentRepository) Delete(id string) error {
	_, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	var comment Comment
	if err := ur.conn.Delete(&comment, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func (ur *CommentRepository) DeleteAllCommentByPostID(postID string) error {

	var comment Comment

	if err := ur.conn.Delete(&comment, "post_id = ?", postID).Error; err != nil {
		return err
	}
	return nil
}

func (ur *CommentRepository) GetCommentCount(postID string) (int, error) {

	var count int64

	if err := ur.conn.Model(&Comment{}).Where("post_id = ?", postID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
