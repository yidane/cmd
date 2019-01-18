package internal

import (
	"bytes"
	"container/list"
	"fmt"
	"github.com/yidane/cmd/opt"
	"strings"
	"unicode"
)

type DefaultParse struct {
}

func (DefaultParse) ParseArg(arg string) (args []*opt.Arg, err error) {
	if len(arg) == 0 {
		args = make([]*opt.Arg, 0)
		return
	}

	l := readStringAsList(arg)
	k, v := "", ""
	node := l.Front()
	for node != nil {
		arg := node.Value.(string)

		if arg[0] != '-' {
			v = arg
			args = append(args, opt.NewArg(len(args), k, v))
			k, v = "", ""
			node = node.Next()
			continue
		}

		if arg[1] == '-' {
			if len(arg) == 1 {
				err = fmt.Errorf("bad flag syntax: %s", arg)
				return
			}
			arg = arg[2:] //set --arg as -arg
		}

		if strings.Contains(arg, "=") { //if container "=",arg and value is here,otherwise the value is next arg
			kv := strings.Split(arg, "=")
			k = kv[0]
			v = strings.Join(kv[1:], "=")

			args = append(args, opt.NewArg(len(args), k, v))
			k, v = "", ""
			node = node.Next()
			continue
		}

		k = arg
		node = node.Next()
		if node == nil {
			args = append(args, opt.NewArg(len(args), k, v))
			continue //end for
		}

		for node != nil {
			arg = node.Value.(string)
			if arg[0] == '-' { //another k
				args = append(args, opt.NewArg(len(args), k, v))
				k, v = "", ""
				break
			}

			//one key match many value
			v = arg
			args = append(args, opt.NewArg(len(args), k, v))
			node = node.Next()
		}
	}

	return
}

func readStringAsList(arg string) (l list.List) {
	l = list.List{}
	buf := bytes.Buffer{}
	for _, v := range strings.TrimSpace(arg) {
		if unicode.IsSpace(v) {
			if buf.Len() > 0 {
				l.PushBack(buf.String())
				buf.Reset()
			}
		} else {
			buf.WriteRune(v)
		}
	}
	if buf.Len() > 0 {
		l.PushBack(buf.String())
		buf.Reset()
	}

	return
}
