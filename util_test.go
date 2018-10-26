package main

import (
	"testing"
)

func Test_getProgramName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		args string
		want string
	}{
		{args: "./cmd", want: "cmd"},
		{args: "./../cmd", want: "cmd"},
		{args: "cmd", want: "cmd"},
		{args: "cmd.exe", want: "cmd.exe"},
		{args: "c:/file/cmd.exe", want: "cmd.exe"},
		{args: "c:\\file\\cmd.exe", want: "cmd.exe"},
	}
	for _, tt := range tests {
		t.Run(tt.args, func(t *testing.T) {
			if got := getProgramName(tt.args); got != tt.want {
				t.Errorf("getProgramName() = %v, want %v", got, tt.want)
			}
		})
	}
}
