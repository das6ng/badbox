package cmd

import (
	"fmt"
	"path"

	"github.com/urfave/cli/v2"
)

var Dirname = &cli.Command{
	Name:   "dirname",
	Usage:  "show the path's directory",
	Action: cmdDirname,
}

func cmdDirname(cCtx *cli.Context) error {
	orig := cCtx.Args().Get(0)
	fmt.Println(path.Dir(orig))
	return nil
}
