package cmd

import (
	"bufio"
	"fmt"
	"github.com/yidane/cmd/internal"
	"github.com/yidane/cmd/opt"
	"github.com/yidane/cmd/pt"
	"io"
	"os"
	"strings"
)

const prefix = " > "

type Context struct {
	option     *opt.ContextOption
	reader     *bufio.Reader
	lastError  error
	BeforeExit func()
}

func NewContext() *Context {
	return &Context{
		option: opt.NewContextOption(),
	}
}

func (ctx *Context) Wait() {
	if ctx.option.Running() {
		pt.Print(prefix)
	}
}

func (ctx *Context) Stop() {
	ctx.option.Stop()
}

func (ctx *Context) SetProgramName(name string) {
	if strings.TrimSpace(name) != "" {
		ctx.option.ProgramName = name
	}
}

func (ctx *Context) GetProgramName() string {
	return ctx.option.ProgramName
}

func (ctx *Context) reset() {
	ctx.option.Args = []*opt.Arg{}
	ctx.option.CommandName = ""
}

func (ctx *Context) recover() {
	err := recover()
	switch err.(type) {
	case nil:
		return
	case error:
		ctx.lastError = err.(error)
	default:
		panic(err) //who is it
	}
}

func (ctx *Context) beforeExit() {
	if ctx.BeforeExit != nil {
		ctx.BeforeExit()
	}
}

func (ctx *Context) exec() {
	if ctx.lastError != nil {
		return
	}

	defer ctx.recover()

	command, exists := internal.GetCommand(ctx.option.CommandName)
	if !exists {
		ctx.lastError = fmt.Errorf("command %s missed", ctx.option.CommandName)
		return
	}

	ctx.option.Args, ctx.lastError = command.ParseArg(ctx.option.CommandString)
	if ctx.lastError != nil {
		return
	}
	ctx.lastError = command.Exec(ctx.option)
}

func (ctx *Context) readCommand() bool {
	ctx.option.CommandName = ""
	ctx.option.CommandString = ""

	if ctx.lastError != nil {
		return false
	}

	//if only run once,the reader is nil
	if ctx.reader == nil {
		if len(os.Args) > 1 {
			cmdName := os.Args[1]
			if _, exists := internal.GetCommand(cmdName); exists {
				ctx.option.CommandName = os.Args[1]
			} else {
				if _, exists := internal.GetCommand(""); !exists {
					ctx.lastError = fmt.Errorf("command %s missed and space command missed", cmdName)
				}
			}
		}
		ctx.option.CommandString = strings.Join(os.Args[1:], " ")
		return true
	}

	input, _, err := ctx.reader.ReadLine()

	switch err {
	case nil, io.EOF:
	default:
		ctx.lastError = err
		return false
	}

	inputString := string(input)
	if len(strings.TrimSpace(inputString)) == 0 {
		return false
	}

	if spaceIndex := strings.Index(inputString, " "); spaceIndex > 0 {
		ctx.option.CommandName = inputString[:spaceIndex]
		ctx.option.CommandString = inputString[spaceIndex:]
	} else {
		ctx.option.CommandName = inputString
	}

	return true
}

func (ctx *Context) handError() {
	if ctx.lastError != nil {
		pt.Error(ctx.lastError)
		ctx.lastError = nil
	}
}

func (ctx *Context) Start() error {
	defer ctx.recover()
	defer ctx.beforeExit()

	ctx.option.Start()
	ctx.reader = bufio.NewReader(os.Stdin)
	ctx.Wait()
	for ctx.option.Running() {
		if ctx.readCommand() {
			ctx.exec()
		}

		ctx.handError()
		ctx.Wait()
	}

	return ctx.lastError
}

func (ctx *Context) Once() error {
	if ctx.readCommand() {
		ctx.exec()
	}

	return ctx.lastError
}
