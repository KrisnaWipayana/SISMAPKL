package auth_handler

import (
	"sismapkl/internal/pkg/schemas"
	"time"

	"github.com/gofiber/fiber/v2"
)

func (h *handler) HandleLoginMahasiswa(c *fiber.Ctx) error {

	input := new(schemas.LoginInputMahasiswa)
	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid input",
		})
	}

	// service login
	loginResponse, err := h.AuthService.LoginMahasiswa(c.Context(), input.NIM, input.Password)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "error",
			"message": "invalid credentials",
			"error":   err.Error(),
		})
	}
	// set cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    loginResponse.Token[len("Bearer "):], // hapus "Bearer " prefix untuk di cookie
		Expires:  time.Now().Add(24 * time.Hour),       // nyesuaiin expires time JWT
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		Path:     "/",
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"status": "success login",
		"token":  loginResponse.Token,
	})
}
