package main

import (
	"fmt"
	"os"
)

//Green change the output's color in terminal
func Green(v ...interface{}) string {
	return fmt.Sprintf("%c[1;0;32m%s%c[0m", 0x1B, v, 0x1B)
}

//Red change the output's color in terminal
func Red(v ...interface{}) string {
	return fmt.Sprintf("%c[1;0;31m%s%c[0m", 0x1B, v, 0x1B)
}

//Yellow change the output's color in terminal
func Yellow(v ...interface{}) string {
	return fmt.Sprintf("%c[1;0;33m%s%c[0m", 0x1B, v, 0x1B)
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
