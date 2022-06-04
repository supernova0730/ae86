package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

func Connect(conf Config) (*gorm.DB, error) {
	gormDialector := postgres.New(postgres.Config{
		DSN: conf.BuildDSN(),
	})

	gormLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			Colorful: true,
			LogLevel: logger.Info,
		},
	)

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
		Logger:         gormLogger,
	}

	return gorm.Open(gormDialector, gormConfig)
}
