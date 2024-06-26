package models

import (
	"github.com/iqbaltc13/back-end-learner-tryout-api/database"

	"time"

	"gorm.io/gorm"
)

type MasterTipeSoalDetail struct {
	gorm.Model
	ID               string `gorm:"primaryKey"`
	Name             string `gorm:"size:1000;not null;column:name" json:"name"`
	MasterTipeSoalid string `gorm:"null;column:master_tipe_soal_id" json:"master_tipe_soal_id"`
	MasterSoalId     string `gorm:"null;column:master_soal_id" json:"master_soal_id"`
	CreatedAt        string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt        string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt        string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`

	MasterSoal MasterSoal `gorm:"foreignKey:ID;references:MasterSoalId"`
}
type MasterSoal struct {
	gorm.Model
	ID               string           `gorm:"primaryKey"`
	Teks             string           `gorm:"null;column:teks" json:"teks"`
	Score            int64            `gorm:"null;column:score" json:"score"`
	Status           string           `gorm:"column:status" json:"status"`
	CreatedAt        string           `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt        string           `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt        string           `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	MasterSoalAssets MasterSoalAssets `gorm:"foreignKey:MasterSoalId;references:ID"`
}
type MasterSoalAssets struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	MasterSoalId string `gorm:"null;column:master_soal_id" json:"master_soal_id"`
	FileID       string `gorm:"null;column:image" json:"image"`
	CreatedAt    string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt    string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt    string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`

	File File `gorm:"foreignKey:FileID"`
}
type MasterSoalPembahasanSoal struct {
	gorm.Model
	ID           string `gorm:"primaryKey"`
	MasterSoalid string `gorm:"null;column:master_soal_id" json:"master_soal_id"`
	Pembahasan   string `gorm:"null;column:pembahasan" json:"pembahasan"`

	CreatedAt                      string                         `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt                      string                         `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt                      string                         `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	MasterSoalPembahasanSoalAssets MasterSoalPembahasanSoalAssets `gorm:"foreignKey:MasterPembahasanSoalId;references:ID"`
}
type MasterSoalPembahasanSoalAssets struct {
	gorm.Model
	ID                     string `gorm:"primaryKey"`
	MasterPembahasanSoalId string `gorm:"null;column:master_pembahasan_soal_id" json:"master_pembahasan_soal_id"`
	FileID                 string `gorm:"null;column:image" json:"image"`
	CreatedAt              string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt              string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt              string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
	File                   File   `gorm:"foreignKey:FileID"`
}

type File struct {
	gorm.Model
	ID                string `gorm:"primaryKey"`
	FullPath          string `gorm:"null;column:full_path" json:"full_path"`
	FullPathThumbnail string `gorm:"null;column:full_path_thumbnail" json:"full_path_thumbnail"`
	CreatedAt         string `gorm:"size:1000;null;column:created_at" json:"created_at"`
	UpdatedAt         string `gorm:"size:1000;null;column:updated_at" json:"updated_at"`
	DeletedAt         string `gorm:"size:1000;null;column:deleted_at" json:"deleted_at"`
}

func (masterTipeSoalDetail *MasterTipeSoalDetail) Save() (*MasterTipeSoalDetail, error) {
	err := database.Database.Create(&masterTipeSoalDetail).Error
	if err != nil {
		return &MasterTipeSoalDetail{}, err
	}
	return masterTipeSoalDetail, nil
}

func (masterTipeSoalDetail *MasterTipeSoalDetail) BeforeSave(*gorm.DB) error {
	currentTime := time.Now()
	masterTipeSoalDetail.CreatedAt = currentTime.Format("2006-01-02 15:04:05")
	return nil
}

func FindMasterTipeSoalDetailById(id string) (MasterTipeSoalDetail, error) {
	var masterTipeSoalDetail MasterTipeSoalDetail
	err := database.Database.Where("id=?", id).Find(&masterTipeSoalDetail).Error
	if err != nil {
		return MasterTipeSoalDetail{}, err
	}
	return masterTipeSoalDetail, nil
}
