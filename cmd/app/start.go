package app

import (
	"github.com/urfave/cli/v2"
)

var startCmd = &cli.Command{
	Name:  "start",
	Usage: "start server",
	Flags: []cli.Flag{
		configFlag,
	},
	Before: loadConfig,
	Action: func(c *cli.Context) error {
		return nil
	},
}
