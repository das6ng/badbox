package main

import (
	"log"
	"os"

	"github.com/das6ng/badbox/cmd"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			cmd.Pwd,
			cmd.Dirname,
			cmd.Which,
			cmd.Exec,
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
