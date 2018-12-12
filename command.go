package cmd

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
)

type Command interface {
	Name() string
	Exec(ctx *Context) error
	Usage() string
}

type CommandParse struct {
	command *Command
	parse   *Parse
}

func (commandParse *CommandParse) ParseArg(arg string) ([]*Arg, error) {
	return (*commandParse.parse).ParseArg(arg)
}

func (commandParse *CommandParse) Name() string {
	return (*commandParse.command).Name()
}

func (commandParse *CommandParse) Exec(ctx *Context) error {
	return (*commandParse.command).Exec(ctx)
}

func (commandParse *CommandParse) Usage() string {
	return (*commandParse.command).Usage()
}

var defaultParse Parse = DefaultParse{}
var (
	commands  = make(map[string]*CommandParse) //TODO:change the data store type so that can query by difference way
	commandMu sync.RWMutex
)

//Register is called when package init
func Register(command Command) {
	if command == nil {
		panic("argument command could not be nil")
	}

	commandMu.Lock()
	defer commandMu.Unlock()

	name := strings.ToLower(command.Name())
	if containCommand(&command) {
		panic(fmt.Errorf("command '%s' contained", name))
	}

	var commandParse CommandParse
	commandParse.command = &command

	if parse, ok := command.(Parse); ok {
		commandParse.parse = &parse
	} else {
		commandParse.parse = &defaultParse
	}

	commands[name] = &commandParse
}

func GetCommand(name string) (commandParse *CommandParse, err error) {
	name = strings.ToLower(name)
	commandParse, ok := commands[name]
	if !ok {
		err = fmt.Errorf("command '%s' not found", name)
	}
	return
}

func containCommand(command *Command) bool {
	if command == nil {
		return false
	}
	name := strings.ToLower((*command).Name())
	storedCommand, err := GetCommand(name)
	if err == nil {
		type0 := reflect.ValueOf(command).Type()
		type1 := reflect.ValueOf(storedCommand).Type()
		return type0 == type1
	}

	return false
}
