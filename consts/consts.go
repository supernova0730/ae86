package consts

import (
	"os"
	"path/filepath"
)

const (
	ConfigEnvPrefix = "ae86"
)

func DefaultConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".ae86/config.yaml")
}
