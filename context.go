package main

import (
	"bufio"
	"fmt"
	"os"
)

type Context struct {
	programName string
	running     bool
	args        *map[string]*Arg
	reader      *bufio.Reader
	lastError   error
	Exiting     func()
}

func NewContext() Context {
	return Context{
		programName: GetProgramName(),
		running:     false,
	}
}

func (context Context) SetProgrameName(name string) {
	context.programName = name
}

func (context Context) GetProgrameName() string {
	if context.programName == "" {
		return "<ProgramName>"
	}

	return context.programName
}

func (context Context) Parse(arg string) {
	if context.lastError != nil {
		return
	}
}

func (context Context) ReadLine() string {
	input, _, err := context.reader.ReadLine()
	if err != nil {
		context.lastError = err
		return ""
	}

	return string(input)
}

func (context Context) Start() {
	context.running = true
	context.reader = bufio.NewReader(os.Stdin)
	for context.running {
		input := context.ReadLine()

		if len(input) == 0 {
			continue
		}

		context.Parse(input)

		fmt.Println(input)
	}

	if context.Exiting != nil {
		context.Exiting()
	}
}

func (context Context) Exec() error {
	//context.Parse(os.Args[1:])
	return nil
}
