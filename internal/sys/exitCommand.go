package sys

import "github.com/yidane/cmd/opt"

type ExitCommand struct {
}

func (ExitCommand) Name() string {
	return "exit"
}

func (ExitCommand) Exec(ctx *opt.ContextOption) error {
	ctx.Stop()
	return nil
}

func (ExitCommand) Usage() string {
	return "type exit to exit program"
}