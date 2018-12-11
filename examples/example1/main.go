package main

import (
	"fmt"
	"github.com/yidane/cmd"
)

type AddCommand struct {
}

func (AddCommand) Name() string {
	return "Add"
}

func (AddCommand) Exec(ctx *cmd.Context) error {
	if len(ctx.Args) == 0 {
		return fmt.Errorf("command Add need args")
	}

	var sum int64
	for _, arg := range ctx.Args {
		fmt.Println(arg)
		i, err := arg.Int64()
		if err != nil {
			return err
		}
		sum += i
	}

	fmt.Println("Sum is ", sum)
	return nil
}

func (AddCommand) Usage() string {
	return "add func"
}

var (
	context *cmd.Context
)

func init() {
	fmt.Println("init")
	cmd.Register(AddCommand{})
	context = cmd.NewContext()
	context.BeforeExit = func() {
		fmt.Println("before exit")
	}
}

func main() {
	context.Start()
}
