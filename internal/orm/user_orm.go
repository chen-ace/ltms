package orm

import (
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string
	Password string
}

type UserOrm struct {
	db *gorm.DB
}

func NewUserOrm(db *gorm.DB) *UserOrm {
	return &UserOrm{db: db}
}

func (userOrm *UserOrm) GetByUsername(username string) (*User, error) {
	var user User
	r := userOrm.db.First(&user, "username = ? and deleted_at IS NULL", username)
	return &user, r.Error
}

func (userOrm *UserOrm) GetById(id int) (*User, error) {
	var user User
	r := userOrm.db.First(&user, id)
	return &user, r.Error
}

func (userOrm *UserOrm) Create(user *User) (uint, error) {
	result := userOrm.db.Create(user)
	if result.Error != nil {
		log.Fatalln(result.Error)
		return 0, result.Error
	}
	return user.ID, nil
}
