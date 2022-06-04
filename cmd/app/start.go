package app

import (
	"ae86/config"
	"ae86/consts"
	"ae86/pkg/logger"
	"github.com/urfave/cli/v2"
)

var startCmd = &cli.Command{
	Name:  "start",
	Usage: "start server",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "config",
			Value: consts.DefaultConfigPath(),
			Usage: "filepath to config.yaml",
		},
	},
	Action: func(c *cli.Context) error {
		configPath := c.String("config")

		conf, err := config.Get(configPath, consts.ConfigEnvPrefix)
		if err != nil {
			return err
		}

		logger.Log.Info("configs initialized...")
		logger.Log.Infof("%#v", conf)
		return nil
	},
}
