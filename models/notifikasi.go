package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"gorm.io/gorm"
)

type Notifikasi struct {
	gorm.Model
	Id          string `gorm:"size:1000;not null;column:id" json:"id"`
	Title       string `gorm:"size:1000;not null;column:title" json:"title"`
	Subtitle    string `gorm:"size:1000;not null;column:subtitle" json:"subtitle"`
	Action      string `gorm:"size:1000;not null;column:action" json:"action"`
	Value       string `gorm:"size:1000;not null;column:value" json:"value"`
	SenderId    string `gorm:"size:1000;not null;column:sender_id" json:"sender_id"`
	ReceiverId  string `gorm:"size:1000;not null;column:receiver_id" json:"receiver_id"`
	CreatedById string `gorm:"size:1000;not null;column:created_at_by_id" json:"current_apk_version_name"`

	ReadAt    string `gorm:"size:1000;not null;column:read_at" json:"read_at"`
	DeletedAt string `gorm:"size:1000;not null;column:deleted_at" json:"deleted_at"`
	CreatedAt string `gorm:"size:1000;not null;column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"size:1000;not null;column:updated_at" json:"updated_at"`

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
