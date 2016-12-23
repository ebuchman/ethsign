package main

import (
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "ethsign"
	app.Usage = ""
	app.Version = "0.0.1"
	app.Author = "Ethan Buchman"
	app.Email = "ethan@coinculture.info"

	app.Flags = []cli.Flag{}

	app.Commands = []cli.Command{
		signCmd,
	}
	app.Run(os.Args)
}

var (
	signCmd = cli.Command{
		Name:   "sign",
		Usage:  "Sign an ethereum transaction",
		Action: cliSign,
		Flags: []cli.Flag{
			keyDirFlag,
			fromFlag,
			toFlag,
			nonceFlag,
			amountFlag,
			gasFlag,
			gasPriceFlag,
			dataFlag,
		},
	}

	keyDirFlag = cli.StringFlag{
		Name:  "keydir",
		Usage: "Path to key dir",
	}

	fromFlag = cli.StringFlag{
		Name:  "from",
		Usage: "From address (hex)",
	}

	toFlag = cli.StringFlag{
		Name:  "to",
		Usage: "To address (hex)",
	}

	amountFlag = cli.IntFlag{
		Name:  "amount",
		Usage: "Number of finney to send (milliether)",
		Value: 0,
	}

	gasFlag = cli.IntFlag{
		Name:  "gas",
		Usage: "Maximum amount of gas for the tx/call",
		Value: 21000,
	}

	gasPriceFlag = cli.IntFlag{
		Name:  "price",
		Usage: "Gas Price, in GWei",
		Value: 25,
	}

	dataFlag = cli.StringFlag{
		Name:  "data",
		Usage: "Data to send to a contract (hex)",
	}

	nonceFlag = cli.IntFlag{
		Name:  "nonce",
		Usage: "Sequence number of the account",
	}
)
