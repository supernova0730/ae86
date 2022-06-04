package transport

import (
	"ae86/internal/container"
	"ae86/internal/transport/rest"
)

func Start(conf rest.Config, restContainer *container.RestContainer) error {
	// telegram bot start

	err := rest.Start(conf, restContainer)
	if err != nil {
		return err
	}

	return nil
}
