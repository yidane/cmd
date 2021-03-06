package sys

import (
	"github.com/yidane/cmd/internal"
	"github.com/yidane/cmd/opt"
	"log"
)

type HelpCommand struct {
}

func (HelpCommand) Name() string {
	return "help"
}

func (HelpCommand) Exec(ctx *opt.ContextOption) error {
	if len(ctx.Args) != 0 {
	}

	internal.EachCommand(func(c *opt.Command) {
		log.Println("  ", (*c).Usage())
	})

	return nil
}

func (HelpCommand) Usage() string {
	return "help or help command"
}
