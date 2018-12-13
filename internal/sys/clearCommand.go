package sys

import (
	"github.com/yidane/cmd/opt"
	"os"
	"os/exec"
)

type ClearCommand struct {
}

func (c *ClearCommand) Name() string {
	return "Clear"
}

func (c *ClearCommand) Exec(ctx *opt.ContextOption) error {
	var cmd *exec.Cmd
	switch ctx.OS {
	case "windows":
		cmd = exec.Command("clear")
	default:
		cmd = exec.Command("cls")
	}

	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (c *ClearCommand) Usage() string {
	return "type Clear to clear the console"
}