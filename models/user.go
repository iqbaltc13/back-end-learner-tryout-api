package models

import (
	"html"
	"strings"

	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id                       string `gorm:"size:1000;not null;" json:"-"`
	name                     string `gorm:"size:1000;not null;" json:"-"`
	username                 string `gorm:"size:1000;not null;unique" json:"-"`
	email                    string `gorm:"size:1000;not null;unique" json:"email"`
	password                 string `gorm:"size:1000;not null;" json:"-"`
	phone                    string `gorm:"size:1000;not null;" json:"-"`
	token_login_mobile       string `gorm:"size:1000;not null;" json:"-"`
	current_apk_version_name string `gorm:"size:1000;not null;" json:"-"`
	current_apk_version_code string `gorm:"size:1000;not null;" json:"-"`
	email_verified_at        string `gorm:"size:1000;not null;" json:"-"`
	email_device_info        string `gorm:"size:1000;not null;" json:"-"`
	email_deleted_at         string `gorm:"size:1000;not null;" json:"-"`

	Entries []Entry
}

func (user *User) Save() (*User, error) {
	err := database.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.password = string(passwordHash)
	user.email = html.EscapeString(strings.TrimSpace(user.email))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := database.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
func FindUserByEmail(email string) (User, error) {
	var user User
	err := database.Database.Where("email=?", email).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id string) (User, error) {
	var user User
	err := database.Database.Preload("Entries").Where("id=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}
