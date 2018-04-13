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
				Usage:  "",
				Action: xuclient.GetInfo,
			},
			{
				Name:   "placeorder",
				Usage:  "",
				Action: xuclient.PlaceOrder,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "lang",
						Value: "english",
						Usage: "",
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
