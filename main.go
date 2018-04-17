package main

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/urfave/cli"
	"time"
)

var VERSION = "v0.0.0-dev"

func main() {
	app := cli.NewApp()
	app.Name = "deanza-devops"
	app.Version = VERSION
	app.Usage = "You need help!"
	app.Action = func(c *cli.Context) error {
		for {
			logrus.Info("Demo gods, have some mercy and let the demo run well")
			time.Sleep(1 * time.Second)
		}
	}

	app.Run(os.Args)
}
