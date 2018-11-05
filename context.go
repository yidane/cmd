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
	programName string
	commandName string
	running     bool
	Args        map[string]*Arg
	reader      *bufio.Reader
	lastError   error
	BeforeExit  func()
}

func NewContext() *Context {
	return &Context{
		programName: os.Args[0],
		running:     false,
		Args:        make(map[string]*Arg),
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

func (ctx *Context) resetArgs() {
	ctx.Args = make(map[string]*Arg)
	ctx.commandName = ""
}

func (ctx *Context) parse(args []string) {
	if ctx.lastError != nil {
		return
	}

	ctx.resetArgs()

	if len(args) == 0 {
		ctx.lastError = fmt.Errorf("args is illegal")
		return
	}

	arg0 := args[0]
	if len(arg0) == 0 || arg0[0] == '-' { //first is command name which can not begin with '-'
		ctx.lastError = fmt.Errorf("command '%s' is illegal that can not begin with '-'", arg0)
		return
	}
	ctx.commandName = strings.TrimSpace(arg0) //set command name

	for i := 1; i < len(args); i++ {
		if ctx.lastError != nil {
			return
		}
		arg := strings.TrimSpace(args[i])
		switch len(arg) {
		case 0:
			continue
		case 1:
			ctx.lastError = fmt.Errorf("bad flag syntax: %s", arg)
			return
		default:
			if arg[0] != '-' {
				ctx.lastError = fmt.Errorf("bad flag syntax: %s", arg)
				return
			}
			if arg[1] == '-' {
				arg = arg[2:]
			}
		}

		k, v := "", ""
		if strings.Contains(arg, "=") { //if container "=",arg and value is here,otherwise the value is next arg
			kv := strings.SplitN(arg, "=", 2)
			k = kv[0]
			v = kv[1]
		} else {
			k = arg
			if i == len(args)-1 { //if it is the last one,it could be bool and set is true
				v = "true"
			} else {
				arg1 := strings.TrimSpace(args[i+1])
				if arg1[0] == '-' {
					v = "true"
				} else {
					v = arg1
					i++
				}
			}
		}

		if _, ok := ctx.Args[k]; ok {
			ctx.lastError = fmt.Errorf("")
			return
		}
		newArg := Arg(v)
		ctx.Args[k] = &newArg
	}
}

func (ctx *Context) readLine() string {
	input, _, err := ctx.reader.ReadLine()

	switch err {
	case nil, io.EOF:
		return strings.ToLower(strings.TrimSpace(string(input)))
	default:
		ctx.lastError = err
		return ""
	}
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

	command, exists := fetchCommand(ctx.commandName)
	if !exists {
		ctx.lastError = fmt.Errorf("'%s' is not a command", ctx.commandName)
		return
	}

	ctx.lastError = (*command).Exec(ctx)
}

func (ctx *Context) Start() error {
	defer ctx.recover()
	defer ctx.beforeExit()

	ctx.running = true
	ctx.reader = bufio.NewReader(os.Stdin)
	fmt.Print(prefix)
	for ctx.running {
		input := ctx.readLine()

		if len(input) == 0 {
			continue
		}

		ctx.parse(strings.Split(input, " "))
		ctx.exec()

		if ctx.lastError != nil {
			fmt.Println(ctx.lastError)
			ctx.lastError = nil
		}

		fmt.Print(prefix)
	}

	return nil
}

func (ctx *Context) Once() error {
	ctx.parse(os.Args[1:])
	ctx.exec()

	if ctx.lastError != nil {
		fmt.Println(ctx.lastError)
	}
	return ctx.lastError
}
