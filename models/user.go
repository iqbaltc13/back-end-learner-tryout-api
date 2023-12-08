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
	Id                    string `gorm:"size:1000;not null;column:id" json:"id"`
	Name                  string `gorm:"size:1000;not null;column:name" json:"name"`
	Username              string `gorm:"size:1000;not null;unique;column:username" json:"username"`
	Email                 string `gorm:"size:1000;not null;unique;column:email" json:"email"`
	Password              string `gorm:"size:1000;not null;column:password" json:"password"`
	Phone                 string `gorm:"size:1000;not null;column:phone" json:"phone"`
	TokenLoginMobile      string `gorm:"size:1000;not null;column:token_login_mobile" json:"token_login_mobile"`
	CurrentApkVersionName string `gorm:"size:1000;not null;column:current_apk_version_name" json:"current_apk_version_name"`
	CurrentApkVersionCode string `gorm:"size:1000;not null;column:current_apk_version_code" json:"current_apk_version_code"`
	VerifiedAt            string `gorm:"size:1000;not null;column:email_verified_at" json:"email_verified_at"`
	DeviceInfo            string `gorm:"size:1000;not null;column:device_info" json:"device_info"`
	DeletedAt             string `gorm:"size:1000;not null;column:deleted_at" json:"deleted_at"`
	CreatedAt             string `gorm:"size:1000;not null;column:created_at" json:"created_at"`
	UpdatedAt             string `gorm:"size:1000;not null;column:updated_at" json:"updated_at"`

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
func isEmailTaken(email string) (error, User) {
	var user User
	err := database.Database.Where("email=?", email).Find(&user).Error
	if err != nil {
		return err, User{}
	}
	return nil, user
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
