package models

import (
	"errors"

	"github.com/SyydMR/To-Do-List/utils"
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"password"`
    Items    []Item `gorm:"foreignKey:UserID" json:"items"`
}
func GetAllUser() []User {
	var users []User
	db.Find(&users)
	return users
}

func GetUserById(id int64) (*User, error) {
    var getUser User
    if err := db.Preload("Items").Where("ID = ?", id).First(&getUser).Error; err != nil {
        return nil, err
    }
    return &getUser, nil
}



func (u *User) Register() (*User, error) {
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}

	u.Password = hashedPassword

	result := db.Create(u)
	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (u *User) Login(plainPassword string) (string, error) {
	var user User

	result := db.Where("username = ?", u.Username).First(&user)
	if result.Error != nil {
		return "", result.Error
	}

	if !utils.CheckPasswordHash(plainPassword, user.Password) {
		return "", errors.New("incorrect password")
	}

	token, err := utils.GenerateJWT(int64(user.ID))
	if err != nil {
		return "", err
	}

	return token, nil
}