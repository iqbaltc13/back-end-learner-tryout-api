package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"time"

	"gorm.io/gorm"
)

type MasterTipeSoal struct {
	gorm.Model
	ID            string `gorm:"primaryKey"`
	Name          string `gorm:"size:1000;not null;column:name" json:"name"`
	MasterUjianid string `gorm:"null;column:master_ujian_id" json:"master_ujian_id"`
	CreatedAt     string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt     string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt     string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
}

func (masterTipeSoal *MasterTipeSoal) Save() (*MasterTipeSoal, error) {
	err := database.Database.Create(&masterTipeSoal).Error
	if err != nil {
		return &MasterTipeSoal{}, err
	}
	return masterTipeSoal, nil
}

func (masterTipeSoal *MasterTipeSoal) BeforeSave(*gorm.DB) error {
	currentTime := time.Now()
	masterTipeSoal.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	return nil
}

func FindMasterTipeSoalById(id string) (MasterTipeSoal, error) {
	var masterTipeSoal MasterTipeSoal
	err := database.Database.Where("id=?", id).Find(&masterTipeSoal).Error
	if err != nil {
		return MasterTipeSoal{}, err
	}
	return masterTipeSoal, nil
}
func FindMasterTipeSoalByIds(ids []string) ([]MasterTipeSoal, error) {
	var masterTipeSoal []MasterTipeSoal
	err := database.Database.Where(ids).Find(&masterTipeSoal).Error
	if err != nil {
		return masterTipeSoal, err
	}
	return masterTipeSoal, nil
}
func FindMasterTipeSoalByMasterUjianIds(ids []string) ([]MasterTipeSoal, error) {
	var masterTipeSoal []MasterTipeSoal
	err := database.Database.Where("master_ujian_id IN ?", ids).Find(&masterTipeSoal).Error
	if err != nil {
		return masterTipeSoal, err
	}
	return masterTipeSoal, nil
}
func FindMasterTipeSoalByMasterUjianId(id string) (MasterTipeSoal, error) {
	var masterTipeSoal MasterTipeSoal
	err := database.Database.Where("master_ujian_id = ?", id).Find(&masterTipeSoal).Error
	if err != nil {
		return MasterTipeSoal{}, err
	}
	return masterTipeSoal, nil
}
