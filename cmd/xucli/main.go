package main

import (
	"os"

	"github.com/indxcrypto/xud/config"
	"github.com/indxcrypto/xud/lib/xuclient"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	config.Viper = config.InitConfig() //Read Config
	//viper := config.Viper              //expose viper
	app := &cli.App{
		Name:  "xucli",
		Usage: "Exchange Union Client",
		Authors: []*cli.Author{
			{
				Name:  "Balamurali Pandranki",
				Email: "balamurali@live.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:   "getinfo",
				Usage:  "complete a task on the list",
				Action: xuclient.GetInfo,
			},
			{
				Name:   "placeorder",
				Usage:  "add a task to the list",
				Action: xuclient.PlaceOrder,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "lang",
						Value: "english",
						Usage: "language for the greeting",
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
