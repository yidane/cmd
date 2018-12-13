package cmd

import (
	"fmt"
	"strconv"
	"strings"
)

type Arg struct {
	Index int
	Key   string
	Value string
}

func NewArg(index int, k, v string) *Arg {
	arg := new(Arg)
	arg.Index = index
	arg.Key = k
	arg.Value = v
	return arg
}

type Parse interface {
	ParseArg(arg string) (args []*Arg, err error)
}

type DefaultParse struct {
}

func (DefaultParse) ParseArg(arg string) (args []*Arg, err error) {
	if len(arg) == 0 {
		args = make([]*Arg, 0)
		return
	}

	argArr := strings.Split(arg, " ")

	for i := 1; i < len(args); i++ {
		arg := strings.TrimSpace(argArr[i])
		switch len(arg) {
		case 0:
			continue
		case 1:
			err = fmt.Errorf("bad flag syntax: %s", arg)
			return
		default:
			if arg[0] != '-' {
				err = fmt.Errorf("bad flag syntax: %s", arg)
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
				arg1 := strings.TrimSpace(argArr[i+1])
				if arg1[0] == '-' {
					v = "true"
				} else {
					v = arg1
					i++
				}
			}
		}

		newArg := NewArg(0, k, v)
		args = append(args, newArg)
	}

	return
}

func (arg *Arg) Int() (int, error) {
	return strconv.Atoi(arg.String())
}

func (arg *Arg) Int64() (int64, error) {
	return strconv.ParseInt(arg.String(), 10, 32)
}

func (arg *Arg) Bool() (bool, error) {
	return strconv.ParseBool(strings.ToLower(arg.String()))
}

func (arg *Arg) Float64() (float64, error) {
	return strconv.ParseFloat(arg.String(), 32)
}

func (arg *Arg) String() string {
	return string(arg.Value)
}
