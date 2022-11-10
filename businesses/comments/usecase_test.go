package comments_test

import (
	"errors"
	"mini-project/businesses/comments"
	_commentMock "mini-project/businesses/comments/mocks"
	"mini-project/businesses/posts"
	"mini-project/businesses/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	commentsService comments.Usecase
	commentsRepo    _commentMock.Repository

	commentDomain comments.Domain
	postDomain    posts.Domain
	userDomain    users.Domain
)

func TestMain(m *testing.M) {
	commentsService = comments.NewCommentUsecase(&commentsRepo)

	commentDomain = comments.Domain{
		ID:      "29102910",
		Comment: "Lorem",
		UserID:  userDomain.ID,
		PostID:  postDomain.ID,
	}
	postDomain = posts.Domain{
		ID:       "29102910",
		FileName: "image.png",
		Caption:  "Lorem Ipsum",
		UserID:   userDomain.ID,
	}

	userDomain = users.Domain{
		ID:   "101010",
		Name: "Orang 1",
	}

	m.Run()

}

func TestCreateComment(t *testing.T) {
	t.Run("CreateComment | Success", func(t *testing.T) {
		commentsRepo.On("Create", &commentDomain).Return(commentDomain, nil).Once()

		_, err := commentsService.Create(&commentDomain)
		assert.Nil(t, err)
	})
}
func TestGetByID(t *testing.T) {
	t.Run("GetByID | Success", func(t *testing.T) {
		commentsRepo.On("GetByID", commentDomain.ID).Return(commentDomain, nil).Once()

		_, err := commentsService.GetByID(commentDomain.ID)
		assert.Nil(t, err)
	})
	t.Run("GetByID | Failed", func(t *testing.T) {
		commentsRepo.On("GetByID", commentDomain.ID).Return(comments.Domain{}, errors.New("Not Found")).Once()

		_, err := commentsService.GetByID(commentDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestGetAllComment(t *testing.T) {
	t.Run("GetAll | Success", func(t *testing.T) {
		commentsRepo.On("GetAll", postDomain.ID).Return([]comments.Domain{commentDomain}, nil).Once()

		_, err := commentsService.GetAll(postDomain.ID)
		assert.Nil(t, err)
	})

}
func TestDeleteComment(t *testing.T) {
	t.Run("Delete | Success", func(t *testing.T) {
		commentsRepo.On("Delete", commentDomain.ID).Return(nil).Once()
		err := commentsService.Delete(commentDomain.ID)
		assert.Nil(t, err)
	})
}
func TestDeleteAllCommentByPostID(t *testing.T) {
	t.Run("Delete | Success", func(t *testing.T) {
		commentsRepo.On("DeleteAllCommentByPostID", postDomain.ID).Return(nil).Once()
		err := commentsService.DeleteAllCommentByPostID(postDomain.ID)
		assert.Nil(t, err)
	})
}
