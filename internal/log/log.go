package log

import (
	"fmt"
	"github.com/fatih/color"
	"runtime"
)

func init() {
	color.NoColor = false

	if runtime.GOOS == "windows" {
		log = windowsLog{}
	} else {
		log = linuxLog{}
	}
}

var (
	log Log
)

type Log interface {
	Errorf(format string, a ...interface{})
	Error(err error)
	Succeed(format string, a ...interface{})
	Warn(format string, a ...interface{})
	Println(a ...interface{})
	Print(a ...interface{})
}

func GetLogger() Log {
	return log
}

//log for linux
type windowsLog struct {
}

func (windowsLog) Errorf(format string, a ...interface{}) {
	fmt.Println(color.Output, color.RedString(format, a...))
}

func (windowsLog) Error(err error) {
	fmt.Println(color.Output, color.RedString(err.Error()))
}

func (windowsLog) Succeed(format string, a ...interface{}) {
	fmt.Println(color.Output, color.BlueString(format, a...))
}

func (windowsLog) Warn(format string, a ...interface{}) {
	fmt.Println(color.Output, color.YellowString(format, a))
}

func (windowsLog) Println(a ...interface{}) {
	fmt.Println(a...)
}

func (windowsLog) Print(a ...interface{}) {
	fmt.Print(a...)
}

//log for windows
type linuxLog struct {
}

func (linuxLog) Errorf(format string, a ...interface{}) {
	color.Red(fmt.Sprintf(format, a...))
}

func (linuxLog) Error(err error) {
	color.Red(err.Error())
}

func (linuxLog) Succeed(format string, a ...interface{}) {
	color.Blue(format, a...)
}

func (linuxLog) Warn(format string, a ...interface{}) {
	color.Yellow(format, a...)
}

func (linuxLog) Println(a ...interface{}) {
	fmt.Println(a...)
}

func (linuxLog) Print(a ...interface{}) {
	fmt.Print(a...)
}
