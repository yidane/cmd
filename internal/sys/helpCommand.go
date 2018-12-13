package internal

import (
	"fmt"
)

type HelpCommand struct {
}

func (HelpCommand) Name() string {
	return "help"
}

func (HelpCommand) Exec(ctx *Context) error {
	if len(ctx.Args) != 0 {
	}

	ctx.EachCommand(func(c *Command) {
		fmt.Println("  ", (*c).Usage())
	})

	return nil
}

func (HelpCommand) Usage() string {
	return "help or help command"
}

func init() {
	Register(HelpCommand{})
}
