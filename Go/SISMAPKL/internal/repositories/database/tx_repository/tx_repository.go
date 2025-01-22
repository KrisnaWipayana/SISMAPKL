package tx_repository

import (
	"database/sql"

	"gorm.io/gorm"
)

type Contract interface {
	StartTransaction(fn func(tx *gorm.DB) error, opts ...*sql.TxOptions) error

	Begin(opts ...*sql.TxOptions) *gorm.DB

	Commit() *gorm.DB

	Rollback() *gorm.DB
}

type repository struct {
	DB *gorm.DB
}

func New(db *gorm.DB) Contract {
	return &repository{DB: db}
}

func (r *repository) StartTransaction(fn func(*gorm.DB) error, opts ...*sql.TxOptions) error {
	return r.DB.Transaction(fn, opts...)
}

// Begin implements Contract.
func (r *repository) Begin(opts ...*sql.TxOptions) *gorm.DB {
	panic("unimplemented")
}

// Commit implements Contract.
func (r *repository) Commit() *gorm.DB {
	panic("unimplemented")
}

// Rollback implements Contract.
func (r *repository) Rollback() *gorm.DB {
	panic("unimplemented")
}
