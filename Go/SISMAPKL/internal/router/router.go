package router

import (
	"sismapkl/internal/app"

	"github.com/gofiber/fiber/v2"
)

// Setup pendukung untuk route
func SetupRouter(f *fiber.App, m app.Middlewares, h app.Handlers) {

	// Connection status
	f.Get("/status-check", h.HandleStatus.HandleStatus)

	// Auth area
	f.Post("/login-mahasiswa", h.HandleLogin.HandleLoginMahasiswa)
}
