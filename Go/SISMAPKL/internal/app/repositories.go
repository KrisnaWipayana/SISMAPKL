package app

import (
	"github.com/go-resty/resty/v2"
	"gorm.io/gorm"
)

type Repositories struct {
}

func RegisterRepositories(db *gorm.DB, resty *resty.Client) Repositories {
	return Repositories{

		// inisiasi repository

	}
}
