package arg

import "context"

func CtxWithExtraArgs(ctx context.Context, extraArgs []string) context.Context {
	return context.WithValue(ctx, "--extra-args", extraArgs) //nolint
}

func ExtraArgsFromCtx(ctx context.Context) []string {
	iv := ctx.Value("--extra-args")
	if iv == nil {
		return nil
	}
	if v, ok := iv.([]string); !ok {
		return nil
	} else {
		return v
	}
}

func SeparateArgs(origArgs []string) (args, extraArgs []string) {
	args = origArgs
	for i, arg := range origArgs {
		if arg != "--" {
			continue
		}
		args = origArgs[:i]
		extraArgs = origArgs[i+1:]
		break
	}
	return
}
