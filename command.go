package cmd

import (
	"fmt"
	"reflect"
	"sync"
)

type Command interface {
	Name() string
	Exec(ctx *Context) error
	Usage() string
}

type CommandGroup interface {
	Group() string
}

//TODO:change the data store type so that can query by difference way
var (
	commands  = make(map[string]*Command)
	groups    = make(map[string]string)
	commandMu sync.RWMutex
)

//Register is called when package init
func Register(command Command) {
	commandMu.Lock()
	defer commandMu.Unlock()

	if command == nil {
		panic("argument command can not be nil")
	}

	name := command.Name()
	c, ok := commands[name]
	if ok {
		type0 := reflect.ValueOf(command).Type()
		type1 := reflect.ValueOf(c).Type()
		if type0 != type1 {
			panic(fmt.Sprint("Command ", name, " complected"))
		}
	}

	commands[command.Name()] = &command
}
