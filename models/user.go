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
	ID                    string `gorm:"primaryKey"`
	Name                  string `gorm:"size:1000;not null;column:name" json:"name"`
	Username              string `gorm:"size:1000;not null;unique;column:username" json:"username"`
	Email                 string `gorm:"size:1000;not null;unique;column:email" json:"email"`
	Password              string `gorm:"size:1000;not null;column:password" json:"password"`
	Phone                 string `gorm:"size:1000;not null;column:phone" json:"phone"`
	TokenLoginMobile      string `gorm:"size:1000;not null;column:token_login_mobile" json:"token_login_mobile"`
	CurrentApkVersionName string `gorm:"size:1000;not null;column:current_apk_version_name" json:"current_apk_version_name"`
	CurrentApkVersionCode string `gorm:"size:1000;not null;column:current_apk_version_code" json:"current_apk_version_code"`

	DeviceInfo string `gorm:"size:1000;not null;column:device_info" json:"device_info"`

	CreatedAt string `gorm:"size:1000;null;column:created_at" json:"created_at"`

	Entries []Entry
}

type Notifikasi struct {
	gorm.Model
	ID          string `gorm:"primaryKey"`
	Title       string `gorm:"size:1000;not null;column:title" json:"title"`
	Subtitle    string `gorm:"size:1000;not null;column:subtitle" json:"subtitle"`
	Action      string `gorm:"size:1000;not null;column:action" json:"action"`
	Value       string `gorm:"size:1000;not null;column:value" json:"value"`
	SenderId    string `gorm:"size:1000;not null;column:sender_id" json:"sender_id"`
	ReceiverID  string `gorm:"size:1000;not null;column:receiver_id" json:"receiver_id"`
	CreatedById string `gorm:"size:1000;not null;column:created_at_by_id" json:"current_apk_version_name"`

	ReadAt    string `gorm:"size:1000;not null;column:read_at" json:"read_at"`
	DeletedAt string `gorm:"size:1000;not null;column:deleted_at" json:"deleted_at"`
	CreatedAt string `gorm:"size:1000;not null;column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"size:1000;not null;column:updated_at" json:"updated_at"`
	Receiver  User   `gorm:"foreignKey:ID;references:ReceiverID"`

	Entries []Entry
}

func (notifikasi *Notifikasi) Save() (*Notifikasi, error) {
	err := database.Database.Create(&notifikasi).Error
	if err != nil {
		return &Notifikasi{}, err
	}
	return notifikasi, nil
}

func FindNotifikasiById(id string) (Notifikasi, error) {
	var notifikasi Notifikasi
	err := database.Database.Preload("Entries").Where("id=?", id).Find(&notifikasi).Error
	if err != nil {
		return Notifikasi{}, err
	}
	return notifikasi, nil
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
