package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Perusahaan struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Nama        string         `gorm:"type:varchar(100);not null" json:"nama"`
	Koordinat   string         `gorm:"type:varchar(100);not null" json:"koordinat"`
	Alamat      string         `gorm:"type:varchar(200);not null" json:"alamat"`
	Kabupaten   string         `gorm:"type:varchar(100);not null" json:"kabupaten"`
	Provinsi    string         `gorm:"type:varchar(100);not null" json:"provinsi"`
	IDMahasiswa *uuid.UUID     `gorm:"type:uuid" json:"id_mahasiswa"`
	Mahasiswa   *Mahasiswa     `gorm:"foreignKey:IDMahasiswa" json:"mahasiswa,omitempty"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Perusahaan) TableName() string {
	return "perusahaan"
}
