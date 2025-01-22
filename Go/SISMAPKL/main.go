package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sismapkl/db/postgresql"
	"sismapkl/internal/app"
	"sismapkl/internal/pkg/env"
	"sismapkl/internal/router"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

func main() {

	//Koneksi ke database
	postgresql.Open()
	DB := postgresql.GetDB()

	// Inisiasi repository
	r := app.RegisterRepositories(DB, resty.New())

	// Inisiasi Service
	s := app.RegisterServices(r)

	// Inisiasi Fiber
	f := fiber.New()

	// Inisiasi Middleware
	m := app.RegisterMiddlewares(s)

	// Inisiasi Handler
	h := app.RegisterHandlers(s)

	// Inisiasi Route
	router.SetupRouter(f, m, h)

	// Mengambil port untuk menjalankan server
	go func() {
		port := env.GetWithDefault("PORT", "3000") // Port yang dipakai
		log.Fatalf("error running server: %v", f.Listen(":"+port))
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	fmt.Println("server stop")
	os.Exit(0)
}
