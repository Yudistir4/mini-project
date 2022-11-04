package users

import (
	"errors"
	"mini-project/businesses/users"
	"os"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) users.Repository {
	return &UserRepository{
		conn: conn,
	}
}

func (ur *UserRepository) CreateUser(domain *users.Domain) (users.Domain, error) {
	user := FromDomain(domain)
	user.ID = uuid.New().String()

	if user.UserType == "student" {
		if err := ur.conn.Omit("LecturerID").Create(user).Error; err != nil {
			return users.Domain{}, err
		}
	} else if user.UserType == "lecturer" {
		if err := ur.conn.Omit("StudentID").Create(&user).Error; err != nil {
			return users.Domain{}, err
		}
	} else {
		if err := ur.conn.Omit("StudentID", "LecturerID").Create(user).Error; err != nil {
			return users.Domain{}, err
		}
	}

	ur.conn.Preload("Lecturer").Preload("Student").First(&user)
	return user.ToDomain(), nil
}
func (ur *UserRepository) GetByEmail(domain *users.Domain) (users.Domain, error) {
	var user User
	if err := ur.conn.Preload("Lecturer").Preload("Student").First(&user, "email = ?", domain.Email).Error; err != nil {
		return users.Domain{}, err
	}

	if user.Password != domain.Password {
		return users.Domain{}, errors.New("Invalid Password")
	}
	return user.ToDomain(), nil
}
func (ur *UserRepository) GetByID(id string) (users.Domain, error) {
	var rec User
	if err := ur.conn.Preload("Lecturer").Preload("Student").First(&rec, "id = ?", id).Error; err != nil {

		return users.Domain{}, err
	}

	return rec.ToDomain(), nil
}
func (ur *UserRepository) GetAllUsers(limit int, page int, userType string, name string) ([]users.Domain, error) {
	startIndex := (page - 1) * limit

	var rec []User
	if name != "" && userType != "" {
		ur.conn.Debug().Preload("Lecturer").Preload("Student").Order("created_at desc").
			Limit(limit).Offset(startIndex).
			Where("name LIKE ?", "%"+name+"%").
			Where("user_type = ?", userType).
			Find(&rec)
	} else if name != "" {
		ur.conn.Debug().Preload("Lecturer").Preload("Student").Order("created_at desc").
			Limit(limit).Offset(startIndex).
			Where("name LIKE ?", "%"+name+"%").
			Find(&rec)

	} else {
		ur.conn.Debug().Preload("Lecturer").Preload("Student").Order("created_at desc").
			Limit(limit).Offset(startIndex).
			Find(&rec)
	}

	usersDomain := []users.Domain{}

	for _, user := range rec {
		usersDomain = append(usersDomain, user.ToDomain())
	}

	return usersDomain, nil
}
func (ur *UserRepository) Update(id string, domain *users.Domain) error {
	user, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	updateUser := FromDomain(&user)
	updateUser.Email = domain.Email
	updateUser.Password = domain.Password
	updateUser.UserType = domain.UserType
	updateUser.Name = domain.Name
	updateUser.Bio = domain.Bio
	updateUser.EksternalLink = domain.EksternalLink
	updateUser.Instagram = domain.Instagram
	updateUser.Linkedin = domain.Linkedin
	updateUser.Whatsapp = domain.Whatsapp
	updateUser.StudentID = domain.StudentID

	if err := ur.conn.Omit("LecturerID", "StudentID", "ProfilePic").Save(&updateUser).Error; err != nil {
		return err
	}

	return nil
}
func (ur *UserRepository) Delete(id string) error {

	var user User

	if err := ur.conn.Delete(&user, "id = ?", id).Error; err != nil {
		return err
	}

	return nil
}

func (ur *UserRepository) UpdateProfilePicture(id string, filename string) error {
	user, err := ur.GetByID(id)
	if err != nil {
		return err
	}

	if user.ProfilePic != filename && user.ProfilePic != "" {
		if err := os.Remove("images/" + user.ProfilePic); err != nil {
			return err
		}
	}

	ur.conn.Model(&User{}).Where("id = ?", id).Update("profile_pic", filename)

	return nil
}
