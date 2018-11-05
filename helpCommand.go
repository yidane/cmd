package cmd

import "fmt"

type HelpCommand struct {
}

func (HelpCommand) Name() string {
	return "help"
}

func (HelpCommand) Exec(ctx *Context) error {
	if len(ctx.Args) != 0 {
	}

	for _, command := range commands {
		fmt.Println("  ", (*command).Usage())
	}

	return nil
}

func (HelpCommand) Usage() string {
	return "help or help command"
}

func init() {
	Register(HelpCommand{})
}
