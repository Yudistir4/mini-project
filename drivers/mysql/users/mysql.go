package users

import (
	"mini-project/businesses/users"

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

func (ur *UserRepository) CreateUser(domain *users.Domain) users.Domain {
	user := FromDomain(domain)
	user.ID = uuid.New().String()
	ur.conn.Save(&user)
	ur.conn.Preload("Lecturer").Preload("Student").First(&user)
	return user.ToDomain()
}
func (ur *UserRepository) GetByEmail(domain *users.Domain) users.Domain {
	var user User
	ur.conn.Preload("Lecturer").Preload("Student").First(&user, "email = ?", domain.Email)

	if user.ID == "" {
		return users.Domain{}
	}

	if user.Password != domain.Password {
		return users.Domain{}
	}
	return user.ToDomain()
}
func (ur *UserRepository) GetByID(id string) users.Domain {
	var rec User
	ur.conn.Preload("Lecturer").Preload("Student").First(&rec, "id = ?", id)
	if rec.ID == "" {
		return users.Domain{}
	}
	return rec.ToDomain()
}
func (ur *UserRepository) GetAllUsers(limit, page int) []users.Domain {
	startIndex := (page - 1) * limit

	var rec []User
	ur.conn.Preload("Lecturer").Preload("Student").Order("created_at desc").Limit(limit).Offset(startIndex).Find(&rec)

	usersDomain := []users.Domain{}

	for _, user := range rec {
		usersDomain = append(usersDomain, user.ToDomain())
	}

	return usersDomain
}
func (ur *UserRepository) Update(id string, domain *users.Domain) users.Domain {
	user := ur.GetByID(id)
	if user.ID == "" {
		return users.Domain{}
	}

	updateUser := FromDomain(&user)
	updateUser.Email = domain.Email
	updateUser.Password = domain.Password
	updateUser.UserType = domain.UserType
	updateUser.Name = domain.Name
	updateUser.ProfilePic = domain.ProfilePic
	updateUser.Bio = domain.Bio
	updateUser.EksternalLink = domain.EksternalLink
	updateUser.Instagram = domain.Instagram
	updateUser.Linkedin = domain.Linkedin
	updateUser.Whatsapp = domain.Whatsapp
	updateUser.StudentID = domain.StudentID

	ur.conn.Save(&updateUser)

	return updateUser.ToDomain()
}
func (ur *UserRepository) Delete(id string) bool {

	userData := ur.GetByID(id)

	if userData.ID == "" {
		return false
	}

	var user User

	if err := ur.conn.Delete(&user, "id = ?", id).Error; err != nil {
		return false
	}

	return true
}
