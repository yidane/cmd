package opt

import (
	"runtime"
)

type Command interface {
	Name() string
	Exec(ctx *ContextOption) error
	Usage() string
}

type Parse interface {
	ParseArg(arg string) (args []*Arg, err error)
}

type ContextOption struct {
	OS            string
	ProgramName   string
	CommandName   string
	running       bool
	Args          []*Arg
	CommandString string
}

func NewContextOption() *ContextOption {
	return &ContextOption{
		OS: runtime.GOOS,
	}
}

func (option *ContextOption) Start() {
	option.running = true
}

func (option *ContextOption) Stop() {
	option.running = false
}

func (option *ContextOption) Running() bool {
	return option.running
}
