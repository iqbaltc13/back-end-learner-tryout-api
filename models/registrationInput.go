package models

type RegistrationInput struct {
	Name                  string `gorm:"size:1000;not null;" json:"name"`
	Username              string `gorm:"size:1000;not null;unique" json:"username"`
	Email                 string `gorm:"size:1000;not null;unique_index;" json:"email" binding:"required"`
	Password              string `gorm:"size:1000;not null;" json:"password"`
	ConfirmPassword       string `gorm:"size:1000;not null;" json:"confirm_password"`
	Phone                 string `gorm:"size:1000;not null;" json:"phone"`
	TokenLoginMobile      string `gorm:"size:1000;not null;" json:"token_login_mobile"`
	CurrentApkVersionName string `gorm:"size:1000;not null;" json:"current_apk_version_name"`
	CurrentApkVersionCode string `gorm:"size:1000;not null;" json:"current_apk_version_code"`
	EmailVerifiedAt       string `gorm:"size:1000;not null;" json:"email_verified_at"`
	CreatedAt             string `gorm:"size:1000;not null;" json:"created_at"`
	DeviceInfo            string `gorm:"size:1000;not null;" json:"device_info"`
}
