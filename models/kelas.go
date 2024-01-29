package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"time"

	"gorm.io/gorm"
)

type Kelas struct {
	gorm.Model
	ID         string `gorm:"primaryKey"`
	Name       string `gorm:"size:1000;not null;column:name" json:"name"`
	Keterangan string `gorm:"not null;column:keterangan" json:"keterangan"`
	Status     string `gorm:"not null;column:status" json:"status"`
	CreatedAt  string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt  string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt  string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
}

func (kelas *Kelas) Save() (*Kelas, error) {
	err := database.Database.Create(&kelas).Error
	if err != nil {
		return &Kelas{}, err
	}
	return kelas, nil
}

func (kelas *Kelas) BeforeSave(*gorm.DB) error {
	currentTime := time.Now()
	kelas.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	return nil
}

func FindKelasById(id string) (Kelas, error) {
	var kelas Kelas
	err := database.Database.Where("id=?", id).Find(&kelas).Error
	if err != nil {
		return Kelas{}, err
	}
	return kelas, nil
}
func FindKelasByIds(ids []string) (Kelas, error) {
	var kelas Kelas
	err := database.Database.Where(ids).Find(&kelas).Error
	if err != nil {
		return Kelas{}, err
	}
	return kelas, nil
}
