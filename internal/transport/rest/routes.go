package rest

import (
	"ae86/internal/container"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, container *container.RestContainer) {
	v1 := r.Group("/api/v1")
	v1.Get("/ping", func(c *fiber.Ctx) error {
		return c.JSON("pong")
	})
}
