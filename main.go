package main

import (
	"context"
	"log"
	"os"

	"github.com/das6ng/badbox/arg"
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
			cmd.Cp,
		},
	}

	args, extraArgs := arg.SeparateArgs(os.Args)
	ctx := arg.CtxWithExtraArgs(context.Background(), extraArgs)
	if err := app.RunContext(ctx, args); err != nil {
		log.Fatal(err)
	}
}
