package posts_test

import (
	"errors"
	_commentMock "mini-project/businesses/comments/mocks"
	_likeMock "mini-project/businesses/likes/mocks"
	"mini-project/businesses/posts"
	_postMock "mini-project/businesses/posts/mocks"
	_saveMock "mini-project/businesses/saves/mocks"
	"mini-project/businesses/users"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	postsService posts.Usecase
	postsRepo    _postMock.Repository
	likesRepo    _likeMock.Repository
	commentsRepo _commentMock.Repository
	savesRepo    _saveMock.Repository

	postDomain posts.Domain
	userDomain users.Domain
)

func TestMain(m *testing.M) {
	postsService = posts.NewPostUsecase(&postsRepo, &commentsRepo, &likesRepo, &savesRepo)

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

func TestCreatePost(t *testing.T) {
	t.Run("CreatePost | Success", func(t *testing.T) {
		postsRepo.On("Create", &postDomain).Return(postDomain, nil).Once()

		_, err := postsService.Create(&postDomain)
		assert.Nil(t, err)
	})
}
func TestGetByID(t *testing.T) {
	t.Run("GetByID | Success", func(t *testing.T) {
		postsRepo.On("GetByID", postDomain.ID).Return(postDomain, nil).Once()
		commentsRepo.On("GetCommentCount", postDomain.ID).Return(10, nil).Once()
		likesRepo.On("GetLikeCount", postDomain.ID).Return(10, nil).Once()
		savesRepo.On("GetByUserIDAndPostID", userDomain.ID, postDomain.ID).Return(nil).Once()
		likesRepo.On("GetByUserIDAndPostID", userDomain.ID, postDomain.ID).Return(nil).Once()

		_, err := postsService.GetByID(userDomain.ID, postDomain.ID)
		assert.Nil(t, err)
	})
	t.Run("GetByID | Failed", func(t *testing.T) {
		postsRepo.On("GetByID", postDomain.ID).Return(postDomain, errors.New("Not Found")).Once()

		_, err := postsService.GetByID(userDomain.ID, postDomain.ID)
		assert.NotNil(t, err)
	})
}
func TestGetAllPost(t *testing.T) {
	t.Run("GetAll | Success", func(t *testing.T) {
		postsRepo.On("GetAll", "").Return([]posts.Domain{postDomain}, nil).Once()
		commentsRepo.On("GetCommentCount", postDomain.ID).Return(10, nil).Once()
		likesRepo.On("GetLikeCount", postDomain.ID).Return(10, nil).Once()
		savesRepo.On("GetByUserIDAndPostID", userDomain.ID, postDomain.ID).Return(nil).Once()
		likesRepo.On("GetByUserIDAndPostID", userDomain.ID, postDomain.ID).Return(nil).Once()

		_, err := postsService.GetAll(userDomain.ID, "")
		assert.Nil(t, err)
	})

}
func TestUpdatePost(t *testing.T) {
	t.Run("Update | Success", func(t *testing.T) {
		postsRepo.On("Update", postDomain.ID, &postDomain).Return(postDomain, nil).Once()

		_, err := postsService.Update(postDomain.ID, &postDomain)
		assert.Nil(t, err)
	})

}
func TestDeletePost(t *testing.T) {
	t.Run("Delete | Success", func(t *testing.T) {
		postsRepo.On("GetByID", postDomain.ID).Return(postDomain, nil).Once()
		commentsRepo.On("DeleteAllCommentByPostID", postDomain.ID).Return(nil).Once()
		likesRepo.On("DeleteAllLikeByPostID", postDomain.ID).Return(nil).Once()
		savesRepo.On("DeleteAllSaveByPostID", postDomain.ID).Return(nil).Once()
		postsRepo.On("Delete", postDomain.ID).Return(nil).Once()

		err := postsService.Delete(postDomain.ID)
		assert.Nil(t, err)
	})

}
func TestDeleteAllPostByUserID(t *testing.T) {
	t.Run("DeleteAllPostByUserID | Success", func(t *testing.T) {
		postsRepo.On("GetAll", userDomain.ID).Return([]posts.Domain{postDomain}, nil).Once()

		commentsRepo.On("DeleteAllCommentByPostID", postDomain.ID).Return(nil).Once()
		likesRepo.On("DeleteAllLikeByPostID", postDomain.ID).Return(nil).Once()
		savesRepo.On("DeleteAllSaveByPostID", postDomain.ID).Return(nil).Once()

		postsRepo.On("DeleteAllPostByUserID", userDomain.ID).Return(nil).Once()

		err := postsService.DeleteAllPostByUserID(userDomain.ID)
		assert.Nil(t, err)
	})

}
func TestSavePost(t *testing.T) {
	t.Run("Save | Success", func(t *testing.T) {
		savesRepo.On("Create", userDomain.ID, postDomain.ID).Return(nil).Once()
		err := postsService.SavePost(userDomain.ID, postDomain.ID)
		assert.Nil(t, err)
	})

}
func TestUnsavePost(t *testing.T) {
	t.Run("Unsave | Success", func(t *testing.T) {
		savesRepo.On("Delete", userDomain.ID, postDomain.ID).Return(nil).Once()
		err := postsService.UnsavePost(userDomain.ID, postDomain.ID)
		assert.Nil(t, err)
	})

}
func TestLikePost(t *testing.T) {
	t.Run("Like | Success", func(t *testing.T) {
		likesRepo.On("Create", userDomain.ID, postDomain.ID).Return(nil).Once()
		err := postsService.LikePost(userDomain.ID, postDomain.ID)
		assert.Nil(t, err)
	})
}
func TestUnlikePost(t *testing.T) {
	t.Run("Unlike | Success", func(t *testing.T) {
		likesRepo.On("Delete", userDomain.ID, postDomain.ID).Return(nil).Once()
		err := postsService.UnlikePost(userDomain.ID, postDomain.ID)
		assert.Nil(t, err)
	})
}
