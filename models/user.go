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
	ID                    string `gorm:"size:1000;not null;" json:"id"`
	Name                  string `gorm:"size:1000;not null;" json:"name"`
	Username              string `gorm:"size:1000;not null;unique" json:"username"`
	Email                 string `gorm:"size:1000;not null;unique" json:"email"`
	Password              string `gorm:"size:1000;not null;" json:"password"`
	Phone                 string `gorm:"size:1000;not null;" json:"phone"`
	TokenLoginMobile      string `gorm:"size:1000;not null;" json:"token_login_mobile"`
	CurrentApkVersionName string `gorm:"size:1000;not null;" json:"current_apk_version_name"`
	CurrentApkVersionCode string `gorm:"size:1000;not null;" json:"current_apk_version_code"`
	EmailVerifiedAt       string `gorm:"size:1000;not null;" json:"email_verified_at"`
	EmailDeviceInfo       string `gorm:"size:1000;not null;" json:"email_device_info"`
	EmailDeletedAt        string `gorm:"size:1000;not null;" json:"email_deleted_at"`

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

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
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
