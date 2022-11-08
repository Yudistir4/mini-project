package likes

import (
	"errors"
	"mini-project/businesses/likes"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type LikeRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) likes.Repository {
	return &LikeRepository{
		conn: conn,
	}
}

func (r *LikeRepository) Create(userID, postID string) error {
	var dataLikeExist Like
	if err := r.conn.First(&dataLikeExist, "user_id = ? AND post_id = ?", userID, postID).Error; err == nil {
		return errors.New("like already exists")
	}

	like := Like{
		ID:     uuid.New().String(),
		UserID: userID,
		PostID: postID,
	}

	if err := r.conn.Save(&like).Error; err != nil {
		return err
	}

	var dataLike Like
	if err := r.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&dataLike, "id = ?", like.ID).Error; err != nil {
		return err
	}

	return nil
}
func (r *LikeRepository) GetByID(id string) (likes.Domain, error) {
	var like Like
	if err := r.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&like, "id = ?", id).Error; err != nil {
		return likes.Domain{}, err
	}

	return like.ToDomain(), nil
}
func (r *LikeRepository) GetByUserIDAndPostID(userID, postID string) error {
	var like Like
	if err := r.conn.First(&like, "user_id = ? AND post_id = ?", userID, postID).Error; err != nil {
		return err
	}

	return nil
}
func (r *LikeRepository) GetAll(postID string) ([]likes.Domain, error) {
	var rec []Like
	r.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").Find(&rec, "post_id = ?", postID)

	likesDomain := []likes.Domain{}

	for _, like := range rec {
		likesDomain = append(likesDomain, like.ToDomain())
	}

	return likesDomain, nil
}
func (r *LikeRepository) Delete(userID, postID string) error {
	err := r.GetByUserIDAndPostID(userID, postID)
	if err != nil {
		return err
	}

	var like Like
	if err := r.conn.Delete(&like, "user_id = ? AND post_id = ?", userID, postID).Error; err != nil {
		return err
	}
	return nil
}
func (r *LikeRepository) DeleteAllLikeByPostID(postID string) error {

	var like Like

	if err := r.conn.Delete(&like, "post_id = ?", postID).Error; err != nil {
		return err
	}
	return nil
}

func (r *LikeRepository) GetLikeCount(postID string) (int, error) {

	var count int64

	if err := r.conn.Model(&Like{}).Where("post_id = ?", postID).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
