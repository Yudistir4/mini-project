package mysql

import (
	"fmt"
	"log"

	"mini-project/drivers/mysql/comments"
	"mini-project/drivers/mysql/lecturers"
	"mini-project/drivers/mysql/likes"
	"mini-project/drivers/mysql/posts"
	"mini-project/drivers/mysql/saves"
	"mini-project/drivers/mysql/students"
	"mini-project/drivers/mysql/users"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     string
}

func (config *ConfigDB) InitDB() *gorm.DB {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DB_USERNAME,
		config.DB_PASSWORD,
		config.DB_HOST,
		config.DB_PORT,
		config.DB_NAME,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when connecting to the database: %s", err)
	}

	log.Println("connected to the database")

	return db
}

func DBMigrate(db *gorm.DB) {
	db.AutoMigrate(&users.User{}, &students.Student{}, &lecturers.Lecturer{}, &posts.Post{}, &comments.Comment{}, &likes.Like{}, &saves.Save{})
}

// func SeedUser(db *gorm.DB) users.User {
// 	password, _ := bcrypt.GenerateFromPassword([]byte("123123"), bcrypt.DefaultCost)

// 	fakeUser, _ := util.CreateFaker[users.User]()

// 	userRecord := users.User{
// 		Email:    fakeUser.Email,
// 		Password: string(password),
// 	}

// 	if err := db.Create(&userRecord).Error; err != nil {
// 		panic(err)
// 	}

// 	var foundUser users.User

// 	db.Last(&foundUser)

// 	foundUser.Password = "123123"

// 	return foundUser
// }

// func CleanSeeders(db *gorm.DB) {
// 	db.Exec("SET FOREIGN_KEY_CHECKS = 0")

// 	categoryResult := db.Exec("DELETE FROM categories")
// 	itemResult := db.Exec("DELETE FROM notes")
// 	userResult := db.Exec("DELETE FROM users")

// 	var isFailed bool = itemResult.Error != nil || userResult.Error != nil || categoryResult.Error != nil

// 	if isFailed {
// 		panic(errors.New("error when cleaning up seeders"))
// 	}

// 	log.Println("Seeders are cleaned up successfully")
// }
