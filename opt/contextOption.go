package opt

import (
	"runtime"
	"strings"
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

func (option *ContextOption) GetArg(name string) (arg *Arg, flag bool) {
	name = strings.ToLower(name)
	for _, v := range option.Args {
		if v.Key() == name {
			arg = v
			flag = true
			return
		}
	}

	return
}

func (option *ContextOption) GetArgs(name string) (args []*Arg, flag bool) {
	name = strings.ToLower(name)
	args = make([]*Arg, 0)
	for _, v := range option.Args {
		if v.Key() == name {
			args = append(args, v)
			flag = true
		}
	}

	return
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
