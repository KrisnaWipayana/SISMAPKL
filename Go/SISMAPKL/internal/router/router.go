package router

import (
	"sismapkl/internal/app"

	"github.com/gofiber/fiber/v2"
)

// Setup pendukung untuk route
func SetupRouter(f *fiber.App, m app.Middlewares, h app.Handlers) {

	// Route penghubung endpoint
	f.Get("/status-check", h.HandleStatus.HandleStatus)
}
