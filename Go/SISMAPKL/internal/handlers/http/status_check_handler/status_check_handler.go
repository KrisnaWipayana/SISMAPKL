package status_check_handler

import "github.com/gofiber/fiber/v2"

func (h *handler) HandleStatus(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}
