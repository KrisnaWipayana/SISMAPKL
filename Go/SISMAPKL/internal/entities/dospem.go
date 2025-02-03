package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dospem struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Nama        string         `gorm:"type:varchar(100);not null" json:"nama"`
	NIP         string         `gorm:"type:varchar(30);not null" json:"nip"`
	Password    string         `gorm:"type:varchar(255);not null" json:"password"`
	IDMahasiswa *uuid.UUID     `gorm:"type:uuid" json:"id_mahasiswa"`
	Mahasiswa   *Mahasiswa     `gorm:"foreignKey:IDMahasiswa" json:"mahasiswa,omitempty"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Dospem) TableName() string {
	return "dospem"
}
