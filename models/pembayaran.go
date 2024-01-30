package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"gorm.io/gorm"
)

type Pembayaran struct {
	gorm.Model
	ID        string `gorm:"primaryKey"`
	Userid    string `gorm:"not null;column:user_id" json:"user_id"`
	Classid   string `gorm:"not null;column:class_id" json:"class_id"`
	Linkfile  string `gorm:"column:link_file" json:"link_file"`
	Amount    int64  `gorm:"column:amount" json:"amount"`
	Status    string `gorm:"column:status" json:"status"`
	CreatedAt string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
}

func (pembayaran *Pembayaran) Save() (*Pembayaran, error) {
	err := database.Database.Create(&pembayaran).Error
	if err != nil {
		return &Pembayaran{}, err
	}
	return pembayaran, nil
}

func FindPembayaranByUserId(Userid string) (Pembayaran, error) {
	var pembayaran Pembayaran
	err := database.Database.Where("user_id = ?", Userid).Find(&pembayaran).Error
	if err != nil {
		return Pembayaran{}, err
	}
	return pembayaran, nil
}
