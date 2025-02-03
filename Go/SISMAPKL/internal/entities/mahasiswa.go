package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Mahasiswa struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	NIM          string         `gorm:"type:varchar(15);not null" json:"nim"`
	Password     string         `gorm:"type:varchar(255);not null" json:"password"`
	Nama         string         `gorm:"type:varchar(100);not null" json:"nama"`
	Kelas        string         `gorm:"type:varchar(10);not null" json:"kelas"`
	Prodi        string         `gorm:"type:varchar(100);not null" json:"prodi"`
	Jurusan      string         `gorm:"type:varchar(100);not null" json:"jurusan"`
	IDDospem     *uuid.UUID     `gorm:"type:uuid" json:"id_dospem"` // Foreign key
	Dospem       *Dospem        `gorm:"foreignKey:IDDospem" json:"dospem,omitempty"`
	IDPerusahaan *uuid.UUID     `gorm:"type:uuid" json:"id_perusahaan"` // Foreign key
	Perusahaan   *Perusahaan    `gorm:"foreignKey:IDPerusahaan" json:"perusahaan,omitempty"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Mahasiswa) TableName() string {
	return "mahasiswa"
}

type Berkas struct {
	ID          uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Nama        string         `gorm:"type:varchar(100);not null" json:"nama"`
	Status      string         `gorm:"type:varchar(20);not null" json:"status"`
	IDMahasiswa uuid.UUID      `gorm:"type:uuid;not null" json:"id_mahasiswa"` // Foreign key
	Mahasiswa   *Mahasiswa     `gorm:"foreignKey:IDMahasiswa" json:"mahasiswa,omitempty"`
	CreatedAt   time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Berkas) TableName() string {
	return "berkas"
}

type Laporan struct {
	ID           uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Laporan      string         `gorm:"type:varchar(100);not null" json:"laporan"`
	StatusDospem string         `gorm:"type:varchar(20);not null" json:"status_dospem"`
	StatusMentor string         `gorm:"type:varchar(20);not null" json:"status_mentor"`
	IDMahasiswa  uuid.UUID      `gorm:"type:uuid;not null" json:"id_mahasiswa"` // Foreign key
	Mahasiswa    *Mahasiswa     `gorm:"foreignKey:IDMahasiswa" json:"mahasiswa,omitempty"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (Laporan) TableName() string {
	return "laporan"
}
