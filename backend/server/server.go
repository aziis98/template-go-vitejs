package server

import (
	"aziis98.com/template-go-vitejs/backend"
	"github.com/gofiber/fiber/v2"
)

type Dependencies struct {
	DB backend.Database
}

func New(deps Dependencies) *fiber.App {
	app := fiber.New()

	registerRoutes(app, deps.DB)

	return app
}
