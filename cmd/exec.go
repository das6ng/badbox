package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/samber/lo"
	"github.com/urfave/cli/v2"
)

var Exec = &cli.Command{
	Name:   "exec",
	Usage:  "exec a command in specified env, such as extra PATH dir",
	Action: cmdExec,
	Flags: []cli.Flag{
		&cli.StringSliceFlag{Name: "env-var", Aliases: []string{"e"}, Usage: "add env var, in format '-e name=val'"},
		&cli.StringSliceFlag{Name: "env-path", Aliases: []string{"p"}, Usage: "add extra item to PATH"},
		&cli.StringSliceFlag{Name: "env-path-overwrite", Aliases: []string{"po"}, Usage: "overwrite PATH env var"},
		&cli.StringFlag{Name: "work-dir", Aliases: []string{"wd"}, Usage: "set work dir"},
	},
}

func cmdExec(cCtx *cli.Context) error {
	arg := cCtx.Args().Slice()
	if len(arg) == 0 {
		return nil
	}
	bin := arg[0]
	args := arg[1:]
	cmd := exec.CommandContext(cCtx.Context, bin, args...)
	cmd.Stdin = os.Stdin
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	// set PATH env var
	envPath := cCtx.StringSlice("env-path-overwrite")
	extraPath := cCtx.StringSlice("env-path")
	const pe = "PATH="
	if len(envPath) > 0 {
		cmd.Env = lo.Map(cmd.Environ(), func(v string, _ int) string {
			if !strings.HasPrefix(v, pe) {
				return v
			}
			return pe + strings.Join(envPath, ":")
		})
	} else if len(extraPath) > 0 {
		cmd.Env = lo.Map(cmd.Environ(), func(v string, _ int) string {
			if !strings.HasPrefix(v, pe) {
				return v
			}
			r := strings.TrimPrefix(v, pe)
			return pe + strings.Join(extraPath, ":") + ":" + r
		})
	}

	// set working dir
	cwd := cCtx.String("work-dir")
	if cwd != "" {
		cmd.Dir = cwd
	}

	// append extra env var
	envs := cCtx.StringSlice("env-var")
	if len(envs) > 0 {
		cmd.Env = append(cmd.Environ(), envs...)
	}

	// run command
	out, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(out))
	return nil
}
