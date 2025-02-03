package auth_repository

import (
	"context"
	"sismapkl/internal/entities"
	"sismapkl/internal/pkg/db"

	"github.com/rotisserie/eris"
	"gorm.io/gorm"
)

type Contract interface {
	CheckUserMahasiswa(
		ctx context.Context,
		tx *gorm.DB,
		NIM string,
	) (*entities.Mahasiswa, error)
}

type repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Contract {
	return &repository{DB: db}
}

func (r *repository) CheckUserMahasiswa(ctx context.Context, tx *gorm.DB, NIM string) (
	*entities.Mahasiswa, error) {

	query := db.
		Use(tx, r.DB).
		WithContext(ctx).
		Model(&entities.Mahasiswa{}).
		Where("nim = ?", NIM)

	var mahasiswa entities.Mahasiswa

	if result := query.First(&mahasiswa); result.Error != nil {
		return nil, eris.New("mahasiswa tidak ditemukan")
	}

	return &mahasiswa, nil
}

// func (r *repository) ChangeProfile(
// 	ctx context.Context,
// 	tx *gorm.DB,
// 	mahasiswa *entities.Mahasiswa,
// 	updatedColumns []string,
// ) error {

// 	query := db.
// 		Use(tx, r.DB).
// 		WithContext(ctx).
// 		Select(updatedColumns)

// 	if result := query.Updates(mahasiswa); result.Error != nil {
// 		return eris.New(result.Error.Error())
// 	}
// 	return nil
// }
