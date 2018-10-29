package cmd

import (
	"fmt"
	"testing"
)

func TestContext_parse(t *testing.T) {
	type field struct {
		arg   []string
		isNil bool
	}

	fields := []field{
		{arg: []string{}, isNil: false},
		{arg: []string{"help"}, isNil: true},
		{arg: []string{"help", "-v"}, isNil: true},
		{arg: []string{"help", "-v", "string"}, isNil: true},
		{arg: []string{"help", "-v", "-s"}, isNil: true},
		{arg: []string{"help", "-v", "--s"}, isNil: true},
		{arg: []string{"help", "--s"}, isNil: true},
		{arg: []string{"help", ""}, isNil: true},
		{arg: []string{"version", "-s", "asd"}, isNil: true},
	}

	ctx := NewContext()

	for i := 0; i < len(fields); i++ {
		f := fields[i]
		ctx.lastError = nil //set lastError nil
		ctx.parse(f.arg)
		if ctx.lastError == nil != f.isNil {
			t.Fatal(ctx.args, ctx.lastError)
		}
		fmt.Println(ctx.commandName)
	}
}

func TestContext_parseOnce(t *testing.T) {
}
