package posts

import (
	"mini-project/businesses/posts"

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

func (ur *PostRepository) Create(domain *posts.Domain) posts.Domain {
	post := FromDomain(domain)
	post.ID = uuid.New().String()
	ur.conn.Save(&post)

	var dataPost Post
	ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&dataPost, "id = ?", post.ID)

	return dataPost.ToDomain()
}
func (ur *PostRepository) GetByID(id string) posts.Domain {
	var post Post
	ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&post, "id = ?", id)

	if post.ID == "" {
		return posts.Domain{}
	}

	return post.ToDomain()
}
func (ur *PostRepository) GetAll() []posts.Domain {
	var rec []Post
	ur.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").Find(&rec)

	postsDomain := []posts.Domain{}

	for _, post := range rec {
		postsDomain = append(postsDomain, post.ToDomain())
	}

	return postsDomain
}
func (ur *PostRepository) Update(id string, domain *posts.Domain) posts.Domain {
	post := ur.GetByID(id)

	updatePost := FromDomain(&post)
	updatePost.Caption = domain.Caption

	if err := ur.conn.Save(&updatePost).Error; err != nil {
		return posts.Domain{}
	}

	return updatePost.ToDomain()
}
func (ur *PostRepository) Delete(id string) bool {
	var post Post
	if err := ur.conn.Delete(&post, "id = ?", id).Error; err != nil {
		return false
	}
	return true
}
func (ur *PostRepository) DeleteAllPostByUserID(userID string) bool {
	var post Post
	if err := ur.conn.Delete(&post, "user_id = ?", userID).Error; err != nil {
		return false
	}
	return true
}
