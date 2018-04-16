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
		Name:     "xucli",
		HelpName: "xucli",
		Usage:    "Exchange Union Client",
		Version:  "0.0.1-alpha-golang",
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
				Name:   "getorders",
				Usage:  "Get all Orders / given OrderID details",
				Action: xuclient.GetOrders,
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "own",
						Value: true,
						Usage: "Get own orders.",
					},
				},
			},
			{
				Name:   "placeorder",
				Usage:  "",
				Action: xuclient.PlaceOrder,
				Flags: []cli.Flag{
					&cli.Float64Flag{
						Name:  "price",
						Usage: "",
					},
					&cli.Float64Flag{
						Name:  "quantity",
						Usage: "",
					},
				},
			},
			{
				Name:   "connect",
				Usage:  "",
				Action: xuclient.Connect,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "",
						Value: "",
						Usage: "",
					},
				},
			},
			{
				Name:   "tokenswap",
				Usage:  "",
				Action: xuclient.TokenSwap,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "",
						Value: "",
						Usage: "",
					},
				},
			},
			{
				Name:   "shutdown",
				Usage:  "Shutdown Running XUD Instance",
				Action: xuclient.Shutdown,
			},
		},
	}
	app.Run(os.Args)
}
