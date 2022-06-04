package app

import (
	"ae86/config"
	"ae86/consts"
	"ae86/pkg/logger"
	"context"
	"github.com/urfave/cli/v2"
)

const ctxConfigKey = "config"

var configFlag = &cli.StringFlag{
	Name:  "config",
	Value: consts.DefaultConfigPath(),
	Usage: "filepath to config.yaml",
}

func loadConfig(c *cli.Context) error {
	configPath := c.String("config")

	conf, err := config.Get(configPath, consts.ConfigEnvPrefix)
	if err != nil {
		return err
	}

	logger.Log.Info("configs initialized...")
	logger.Log.Infof("%#v", conf)

	c.Context = setConfigTo(c.Context, conf)
	return nil
}

func setConfigTo(ctx context.Context, conf config.Config) context.Context {
	return context.WithValue(ctx, ctxConfigKey, conf)
}

func getConfigFrom(ctx context.Context) config.Config {
	return ctx.Value(ctxConfigKey).(config.Config)
}
