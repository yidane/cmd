package pt

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
)

var x = runtime.GOOS != "windows"

func Errorf(format string, a ...interface{}) {
	if x {
		color.Red(format, a...)
		return
	}

	fmt.Println(color.Output, color.RedString(format, a...))
}

func Error(err error) {
	if x {
		color.Red(err.Error())
		return
	}
	fmt.Println(color.Output, color.RedString(err.Error()))
}

func Succeed(format string, a ...interface{}) {
	if x {
		color.Blue(format, a...)
		return
	}

	fmt.Println(color.Output, color.BlueString(format, a...))
}

func Warn(format string, a ...interface{}) {
	if x {
		color.Yellow(format, a...)
		return
	}

	fmt.Println(color.Output, color.YellowString(format, a))
}

func Println(a ...interface{}) {
	fmt.Println(a...)
}

func Print(a ...interface{}) {
	fmt.Print(a...)
}
