package app

import (
	"ae86/pkg/logger"
	"github.com/urfave/cli/v2"
	"os"
)

func Run() {
	root := &cli.App{
		Name:  "ae86",
		Usage: "delivery service application",
		Commands: []*cli.Command{
			configCmd,
			startCmd,
		},
	}
	if err := root.Run(os.Args); err != nil {
		logger.Log.Fatal(err)
	}
}
