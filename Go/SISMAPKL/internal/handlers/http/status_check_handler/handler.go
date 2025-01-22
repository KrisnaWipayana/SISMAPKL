package status_check_handler

import "github.com/gofiber/fiber/v2"

type Contract interface {
	// 	HandleStatus(c *fiber.Ctx) error
	HandleStatus(c *fiber.Ctx) error
}

type handler struct{}

func New() Contract {
	return &handler{}
}
