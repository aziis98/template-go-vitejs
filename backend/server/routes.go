package server

import (
	"aziis98.com/template-go-vitejs/backend"
	"github.com/gofiber/fiber/v2"
)

func registerRoutes(r fiber.Router, db backend.Database) {
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./out/frontend/index.html")
	})

	r.Static("/", "./out/frontend")
}
