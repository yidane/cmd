package opt

import (
	"strconv"
	"strings"
)

type Arg struct {
	index int
	key   string
	value string
}

func NewArg(index int, k, v string) *Arg {
	arg := new(Arg)
	arg.index = index
	arg.key = strings.ToLower(k)
	arg.value = v
	return arg
}

func (arg *Arg) Index() int {
	return arg.index
}

func (arg *Arg) Key() string {
	return arg.key
}

func (arg *Arg) Value() string {
	return arg.value
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
	return strings.TrimSpace(arg.value)
}
