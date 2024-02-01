package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"time"

	"gorm.io/gorm"
)

type Class struct {
	gorm.Model
	ID           string        `gorm:"primaryKey"`
	Name         string        `gorm:"size:1000;not null;column:name" json:"name"`
	Keterangan   string        `gorm:"not null;column:keterangan" json:"keterangan"`
	Status       string        `gorm:"not null;column:status" json:"status"`
	CreatedAt    string        `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt    string        `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt    string        `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	MasterUjians []MasterUjian `gorm:"foreignKey:ClassID;references:ID"`
}

type MasterUjian struct {
	gorm.Model
	ID                    string `gorm:"primaryKey"`
	Name                  string `gorm:"size:1000;not null;column:name" json:"name"`
	Keterangan            string `gorm:"not null;column:keterangan" json:"keterangan"`
	ClassID               string `gorm:"not null;column:class_id" json:"class_id"`
	MaxPeserta            string `gorm:"not null;column:max_peserta" json:"max_peserta"`
	NilaiPerSoal          string `gorm:"not null;column:nilai_per_soal" json:"nilai_per_soal"`
	IsGroupBySoalCategory string `gorm:"not null;column:is_group_by_soal_category" json:"is_group_by_soal_category"`
	CreatedAt             string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt             string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt             string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	Class                 Class  `gorm:"foreignKey:ClassID;references:"ID`
}

func (class *Class) Save() (*Class, error) {
	err := database.Database.Create(&class).Error
	if err != nil {
		return &Class{}, err
	}
	return class, nil
}

func (class *Class) BeforeSave(*gorm.DB) error {
	currentTime := time.Now()
	class.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	return nil
}

func FindClassById(id string) (Class, error) {
	var class Class
	err := database.Database.Where("id=?", id).Find(&class).Error
	if err != nil {
		return Class{}, err
	}
	return class, nil
}
func FindClassByIds(ids []string) ([]Class, error) {
	var class []Class
	err := database.Database.Where(ids).Find(&class).Error
	if err != nil {
		return class, err
	}
	return class, nil
}

func FindClassAndMasterUjianByIds(ids []string) ([]Class, error) {
	var class []Class
	err := database.Database.Preload("MasterUjians").Where(ids).Find(&class).Error
	if err != nil {
		return class, err
	}
	return class, nil
}
