package cmd

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

var Which = &cli.Command{
	Name:   "which",
	Usage:  "find the excutable in PATH",
	Action: cmdWhich,
}

func cmdWhich(cCtx *cli.Context) error {
	orig := cCtx.Args().Get(0)
	res, _ := exec.LookPath(orig)
	fmt.Println(res)
	return nil
}
