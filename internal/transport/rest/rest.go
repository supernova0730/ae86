package rest

import (
	"ae86/internal/container"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
)

func Start(config Config, container *container.RestContainer) error {
	app := fiber.New(fiber.Config{
		JSONEncoder:           json.Marshal,
		JSONDecoder:           json.Unmarshal,
		DisableStartupMessage: true,
	})

	// set middlewares
	RegisterRoutes(app, container)

	address := config.Address()
	if config.TLSEnable {
		return app.ListenTLS(address, config.CertFile, config.KeyFile)
	}

	return app.Listen(address)
}
