package app

import (
	"ae86/config"
	"ae86/consts"
	"ae86/pkg/logger"
	"github.com/urfave/cli/v2"
)

var configCmd = &cli.Command{
	Name:  "config",
	Usage: "generate initial config",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "path",
			Value: consts.DefaultConfigPath(),
			Usage: "config filepath",
		},
	},
	Action: func(c *cli.Context) error {
		configPath := c.String("path")
		err := config.GenerateDefault(configPath)
		if err != nil {
			return err
		}

		logger.Log.Infof("config file is successfully generated to: %s", configPath)
		return nil
	},
}
