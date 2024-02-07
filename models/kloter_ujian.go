package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"time"

	"gorm.io/gorm"
)

type KloterUjian struct {
	gorm.Model
	ID                 string `gorm:"primaryKey"`
	Name               string `gorm:"size:1000;not null;column:name" json:"name"`
	MasterUjianid      string `gorm:"null;column:master_ujian_id" json:"master_ujian_id"`
	IsWaktuPelaksanaan int16  `gorm:"null;column:is_watu_pelaksanaan" json:"is_waktu_pelaksanaan"`
	DurasiMenit        int64  `gorm:"not null;column:durasi_menit" json:"durasi_menit"`
	StartDatetime      string `gorm:"size:1000;null;column:start_datetime" json:"start_datetime"`
	EndDatetime        string `gorm:"size:1000;null;column:end_datetime" json:"end_datetime"`
	CreatedAt          string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt          string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt          string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
}

func (kloterUjian *KloterUjian) Save() (*KloterUjian, error) {
	err := database.Database.Create(&kloterUjian).Error
	if err != nil {
		return &KloterUjian{}, err
	}
	return kloterUjian, nil
}

func (kloterUjian *KloterUjian) BeforeSave(*gorm.DB) error {
	currentTime := time.Now()
	kloterUjian.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	return nil
}

func FindKloterUjianById(id string) (KloterUjian, error) {
	var kloterUjian KloterUjian
	err := database.Database.Where("id=?", id).Find(&kloterUjian).Error
	if err != nil {
		return KloterUjian{}, err
	}
	return kloterUjian, nil
}
func FindKloterUjianByIds(ids []string) ([]KloterUjian, error) {
	var kloterUjian []KloterUjian
	err := database.Database.Where(ids).Find(&kloterUjian).Error
	if err != nil {
		return kloterUjian, err
	}
	return kloterUjian, nil
}
func FindKloterUjianByMasterUjianIds(ids []string) ([]KloterUjian, error) {
	var kloterUjian []KloterUjian
	err := database.Database.Where("master_ujian_id IN ?", ids).Find(&kloterUjian).Error
	if err != nil {
		return kloterUjian, err
	}
	return kloterUjian, nil
}
func FindKloterUjianByMasterUjianId(id string) (KloterUjian, error) {
	var kloterUjian KloterUjian
	err := database.Database.Where("master_ujian_id = ?", id).Find(&kloterUjian).Error
	if err != nil {
		return KloterUjian{}, err
	}
	return kloterUjian, nil
}
