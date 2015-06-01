package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "autil"
	app.Version = Version
	app.Usage = ""
	app.Author = "Hirokazu Miyaji"
	app.Email = "hirokazu.miyaji@gmail.com"
	app.Commands = Commands

	app.Run(os.Args)
}
