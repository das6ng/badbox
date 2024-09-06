package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var Pwd = &cli.Command{
	Name:    "pwd",
	Aliases: []string{"cwd"},
	Usage:   "show current directory",
	Action:  cmdPwd,
}

func cmdPwd(cCtx *cli.Context) error {
	res, _ := os.Getwd()
	fmt.Println(res)
	return nil
}
