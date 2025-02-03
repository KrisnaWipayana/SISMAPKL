package app

import (
	"sismapkl/internal/repositories/database/auth_repository"
	"sismapkl/internal/repositories/database/tx_repository"

	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type Repositories struct {
	TxRepository   tx_repository.Contract
	AuthRepository auth_repository.Contract
}

func RegisterRepositories(db *gorm.DB, resty *resty.Client) Repositories {
	return Repositories{
		// inisiasi repository
		AuthRepository: auth_repository.New(db),
	}
}
