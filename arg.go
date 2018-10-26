package main

import (
	"strconv"
	"strings"
)

type Arg struct {
	k string
	v string
}

func (arg *Arg) Name() string {
	return arg.k
}

func (arg *Arg) Int() (int, error) {
	return strconv.Atoi(arg.v)
}

func (arg *Arg) Int64() (int64, error) {
	return strconv.ParseInt(arg.v, 10, 32)
}

func (arg *Arg) Bool() (bool, error) {
	return strconv.ParseBool(strings.ToLower(arg.v))
}

func (arg *Arg) Float64() (float64, error) {
	return strconv.ParseFloat(arg.v, 32)
}

func (arg *Arg) String() string {
	return arg.v
}
