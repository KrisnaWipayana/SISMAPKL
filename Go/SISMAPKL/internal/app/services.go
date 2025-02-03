package app

import "sismapkl/internal/services/auth_service"

type Services struct {
	AuthService auth_service.Contract
}

func RegisterServices(r Repositories) Services {
	return Services{
		// inisiasi service
		AuthService: auth_service.New(
			r.TxRepository,
			r.AuthRepository,
		),
	}
}
