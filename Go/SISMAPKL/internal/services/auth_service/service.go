package auth_service

import (
	"context"
	"sismapkl/internal/pkg/schemas"
	"sismapkl/internal/repositories/database/auth_repository"
	"sismapkl/internal/repositories/database/tx_repository"
)

type Contract interface {
	LoginMahasiswa(
		ctx context.Context,
		NIM string,
		Password string,
	) (*schemas.LoginInputMahasiswa, error)
}

type service struct {
	TxRepository   tx_repository.Contract
	AuthRepository auth_repository.Contract
}

func New(
	txRepository tx_repository.Contract,
	authRepository auth_repository.Contract,
) Contract {
	return &service{
		TxRepository:   txRepository,
		AuthRepository: authRepository,
	}
}
