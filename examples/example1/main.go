package main

import (
	"fmt"
	"github.com/yidane/cmd"
)

var (
	context *cmd.Context
)

func init() {
	context = cmd.NewContext()
	context.BeforeExit = func() {
		fmt.Println("before exit")
	}
}

func main() {
	context.Start()
}
