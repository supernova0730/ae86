package postgres

import "fmt"

type Config struct {
	Username string
	Password string
	Host     string
	Port     string
	Database string
	SSLMode  string
}

func (c Config) BuildDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.Username,
		c.Password,
		c.Database,
		c.SSLMode,
	)
}
