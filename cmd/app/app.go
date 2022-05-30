package app

import (
	"ae86/config"
	"ae86/consts"
	"ae86/pkg/logger"
	"flag"
	"os"
	"path/filepath"
)

func Start() {
	logger.Init()

	// todo: generate default config
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

	return defaultConfigPath(), nil
}

func defaultConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".ae86/config.yaml")
}
