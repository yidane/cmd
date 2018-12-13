package main

import (
	"github.com/yidane/cmd"
	"github.com/yidane/cmd/fmt"
	"github.com/yidane/cmd/opt"
	"strings"
)

func init() {
	cmd.Register(&AICommand{})
}

type AICommand struct {
}

func (*AICommand) ParseArg(arg string) (args []*opt.Arg, err error) {
	args = append(args, opt.NewArg(0, "", arg))
	return
}

func (*AICommand) Name() string {
	return "AI"
}

func (*AICommand) Exec(ctx *opt.ContextOption) error {
	if len(ctx.Args) > 0 {
		s := ctx.Args[0].Value
		s = strings.Replace(s, "?", "", -1)
		s = strings.Replace(s, "？", "", -1)

		if strings.HasSuffix(s, "吗") {
			s = strings.Replace(s, "吗", "", -1)
		}

		s = s + "!"

		fmt.Println(s)
	}
	return nil
}

func (*AICommand) Usage() string {
	return "AI"
}

func main() {
	context := cmd.NewContext()
	context.Start()
}
