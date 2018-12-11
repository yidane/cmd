package cmd

import (
	"os"
	"os/exec"
)

/*
 c := exec.Command("cls")
    c.Stdout = os.Stdout
    c.Run()
*/

type ClearCommand struct {
}

func (c *ClearCommand) Name() string {
	return "Clear"
}

//TODO:should support windows and mac
func (c *ClearCommand) Exec(ctx *Context) error {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

func (c *ClearCommand) Usage() string {
	return "type Clear to clear the console"
}

func init() {
	Register(&ClearCommand{})
}
