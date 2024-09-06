package cmd

import (
	"fmt"

	"github.com/otiai10/copy"
	"github.com/urfave/cli/v2"
)

var Cp = &cli.Command{
	Name:   "cp",
	Usage:  "simple copy file",
	Action: cmdCp,
}

func cmdCp(cCtx *cli.Context) error {
	src := cCtx.Args().Get(0)
	dst := cCtx.Args().Get(1)
	if src == "" || dst == "" {
		return fmt.Errorf("src or dst not specified")
	}
	return copy.Copy(src, dst)
}
