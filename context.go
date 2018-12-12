package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const prefix = " > "

type Context struct {
	programName   string
	commandName   string
	running       bool
	Args          []*Arg
	reader        *bufio.Reader
	lastError     error
	commandString string
	BeforeExit    func()
}

func NewContext() *Context {
	return &Context{
		programName: os.Args[0],
		running:     false,
		Args:        []*Arg{},
	}
}

func (ctx *Context) Stop() {
	ctx.running = false
}

func (ctx *Context) EachCommand(f func(command *Command)) {
	for _, c := range commands {
		f(c.command)
	}
}

func (ctx *Context) SetProgramName(name string) {
	if strings.TrimSpace(name) != "" {
		ctx.programName = name
	}
}

func (ctx *Context) GetProgramName() string {
	return ctx.programName
}

func (ctx *Context) reset() {
	ctx.Args = []*Arg{}
	ctx.commandName = ""
}

func (ctx *Context) recover() {
	err := recover()
	switch err.(type) {
	case nil:
		return
	case error:
		ctx.lastError = err.(error)
	default:
		fmt.Println(err)
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

	command, err := GetCommand(ctx.commandName)
	if err != nil {
		ctx.lastError = err
		return
	}

	ctx.Args, ctx.lastError = command.ParseArg(ctx.commandString)
	if ctx.lastError != nil {
		return
	}
	ctx.lastError = command.Exec(ctx)
}

func (ctx *Context) readCommand() bool {
	if ctx.lastError != nil {
		return false
	}

	if ctx.reader == nil {
		ctx.commandString = strings.Join(os.Args[1:], " ")
		return false
	}

	input, _, err := ctx.reader.ReadLine()

	switch err {
	case nil, io.EOF:
	default:
		ctx.lastError = err
		return false
	}

	ctx.commandString = string(input)

	return false
}

func (ctx *Context) Start() error {
	defer ctx.recover()
	defer ctx.beforeExit()

	ctx.running = true
	ctx.reader = bufio.NewReader(os.Stdin)
	fmt.Print(prefix)
	for ctx.running {
		if ctx.readCommand() {
			ctx.exec()
		}

		if ctx.lastError != nil {
			fmt.Println(ctx.lastError)
			ctx.lastError = nil
		}

		fmt.Print(prefix)
	}

	return ctx.lastError
}

func (ctx *Context) Once() error {
	if ctx.readCommand() {
		ctx.exec()
	}

	return ctx.lastError
}
