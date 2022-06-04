package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Connect(conf Config) (*gorm.DB, error) {
	gormDialector := postgres.New(postgres.Config{
		DSN: conf.BuildDSN(),
	})

	gormConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{SingularTable: true},
	}

	return gorm.Open(gormDialector, gormConfig)
}
