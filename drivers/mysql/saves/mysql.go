package saves

import (
	"errors"
	"mini-project/businesses/saves"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SaveRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) saves.Repository {
	return &SaveRepository{
		conn: conn,
	}
}

func (r *SaveRepository) Create(userID, postID string) error {
	var dataSaveExist Save
	if err := r.conn.First(&dataSaveExist, "user_id = ? AND post_id = ?", userID, postID).Error; err == nil {
		return errors.New("save already exists")
	}

	save := Save{
		ID:     uuid.New().String(),
		UserID: userID,
		PostID: postID,
	}

	if err := r.conn.Save(&save).Error; err != nil {
		return err
	}

	var dataSave Save
	if err := r.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&dataSave, "id = ?", save.ID).Error; err != nil {
		return err
	}

	return nil
}
func (r *SaveRepository) GetByID(id string) (saves.Domain, error) {
	var save Save
	if err := r.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").First(&save, "id = ?", id).Error; err != nil {
		return saves.Domain{}, err
	}

	return save.ToDomain(), nil
}
func (r *SaveRepository) GetByUserIDAndPostID(userID, postID string) error {
	var save Save
	if err := r.conn.First(&save, "user_id = ? AND post_id = ?", userID, postID).Error; err != nil {
		return err
	}

	return nil
}
func (r *SaveRepository) GetAll(postID string) ([]saves.Domain, error) {
	var rec []Save
	r.conn.Preload("User").Preload("User.Student").Preload("User.Lecturer").Find(&rec, "post_id = ?", postID)

	savesDomain := []saves.Domain{}

	for _, save := range rec {
		savesDomain = append(savesDomain, save.ToDomain())
	}

	return savesDomain, nil
}
func (r *SaveRepository) Delete(userID, postID string) error {
	err := r.GetByUserIDAndPostID(userID, postID)
	if err != nil {
		return err
	}

	var save Save
	if err := r.conn.Delete(&save, "user_id = ? AND post_id = ?", userID, postID).Error; err != nil {
		return err
	}
	return nil
}
func (r *SaveRepository) DeleteAllSaveByPostID(postID string) error {

	var save Save

	if err := r.conn.Delete(&save, "post_id = ?", postID).Error; err != nil {
		return err
	}
	return nil
}
