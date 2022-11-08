package posts

import (
	"mini-project/businesses/posts"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) posts.Repository {
	return &PostRepository{
		conn: conn,
	}
}

func (ur *PostRepository) Create(domain *posts.Domain) (posts.Domain, error) {
	post := FromDomain(domain)
	post.ID = uuid.New().String()

	if err := ur.conn.Save(&post).Error; err != nil {
		return posts.Domain{}, err
	}

	var dataPost Post
	if err := ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&dataPost, "id = ?", post.ID).Error; err != nil {
		return posts.Domain{}, err
	}

	return dataPost.ToDomain(), nil
}
func (ur *PostRepository) GetByID(id string) (posts.Domain, error) {
	var post Post
	if err := ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&post, "id = ?", id).Error; err != nil {
		return posts.Domain{}, err
	}

	return post.ToDomain(), nil
}
func (ur *PostRepository) GetAll(userID string) ([]posts.Domain, error) {
	var rec []Post
	if userID != "" {
		ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").Find(&rec, "user_id = ?", userID)

	} else {

		ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").Find(&rec)
	}

	postsDomain := []posts.Domain{}

	for _, post := range rec {
		postsDomain = append(postsDomain, post.ToDomain())
	}

	return postsDomain, nil
}
func (ur *PostRepository) Update(id string, domain *posts.Domain) (posts.Domain, error) {
	post, err := ur.GetByID(id)
	if err != nil {
		return posts.Domain{}, err
	}

	updatePost := FromDomain(&post)
	updatePost.Caption = domain.Caption

	if err := ur.conn.Save(&updatePost).Error; err != nil {
		return posts.Domain{}, err
	}
	post.Caption = domain.Caption
	return post, nil
}
func (ur *PostRepository) Delete(id string) error {
	postData, err := ur.GetByID(id)
	if err != nil {
		return err
	}
	if err := os.Remove("images/" + postData.FileName); err != nil {
		return err
	}

	var post Post
	if err := ur.conn.Delete(&post, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func (ur *PostRepository) DeleteAllPostByUserID(userID string) error {
	var posts []Post
	ur.conn.Find(&posts, "user_id = ?", userID)
	for _, val := range posts {
		if err := os.Remove("images/" + val.FileName); err != nil {
			return err
		}
	}

	var post Post

	if err := ur.conn.Delete(&post, "user_id = ?", userID).Error; err != nil {
		return err
	}
	return nil
}
