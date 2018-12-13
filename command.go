package cmd

import (
	"os"

	"github.com/yidane/cmd/internal"
	"github.com/yidane/cmd/internal/sys"
	"github.com/yidane/cmd/opt"
)

//Register is called when package init
func Register(command opt.Command) {
	internal.Register(command)
}

//Register system command
func init() {
	Register(&sys.ClearCommand{})
	Register(&sys.ExitCommand{})
	Register(&sys.HelpCommand{})
}

//maybe needed, maybe not
func EachCommand(f func(command *opt.Command)) {
	internal.EachCommand(f)
}

func getProgramName(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' || name[i] == '\\' {
			return name[i+1:]
		}
	}

	return name
}

func GetProgramName() string {
	name := os.Args[0]

	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' || name[i] == '\\' {
			return name[i+1:]
		}
	}
	return name
}
