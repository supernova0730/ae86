package app

import (
	"ae86/config"
	"ae86/pkg/client/postgres"
	"fmt"
)

func Run(conf config.Config) error {
	db, err := postgres.Connect(postgres.Config{
		Username: conf.DB.Username,
		Password: conf.DB.Password,
		Host:     conf.DB.Host,
		Port:     conf.DB.Port,
		Database: conf.DB.Database,
		SSLMode:  conf.DB.SSLMode,
	})
	if err != nil {
		return err
	}

	fmt.Println(db)
	return nil
}
