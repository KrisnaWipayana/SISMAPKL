package auth_handler

import (
	"sismapkl/internal/services/auth_service"

	"github.com/gofiber/fiber/v2"
)

type Contract interface {
	HandleLoginMahasiswa(c *fiber.Ctx) error
}

type handler struct {
	AuthService auth_service.Contract
}

func New(
	authService auth_service.Contract,
) Contract {
	return &handler{
		AuthService: authService,
	}
}
