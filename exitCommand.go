package main

type ExitCommand struct {
}

func (ExitCommand) Name() string {
	return "exit"
}

func (ExitCommand) Exec(ctx *Context) error {
	ctx.running = false
	return nil
}

func (ExitCommand) Usage() string {
	return "type exit to exit program"
}

func init() {
	Register(ExitCommand{})
}
