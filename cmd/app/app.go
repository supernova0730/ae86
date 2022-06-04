package app

import (
	"ae86/config"
	"ae86/consts"
	"ae86/pkg/logger"
	"flag"
)

func Start() {
	logger.Init()

	configPath, err := getConfigPath()
	if err != nil {
		logger.Log.Fatal(err)
	}

	conf, err := config.Get(configPath, consts.ConfigEnvPrefix)
	if err != nil {
		logger.Log.Fatal(err)
	}

	logger.Log.Info("configs initialized...")
	logger.Log.Infof("%#v", conf)
}

func getConfigPath() (string, error) {
	configPath := flag.String("config", "", "config file")
	flag.Parse()

	if *configPath != "" {
		return *configPath, nil
	}

	return consts.DefaultConfigPath(), nil
}
