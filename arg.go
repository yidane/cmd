package cmd

import (
	"strconv"
	"strings"
)

type Arg string

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
	return string(*arg)
}
