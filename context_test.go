package cmd

import (
	"testing"
)

func TestContext_parse(t *testing.T) {
	ctx := NewContext()
	ctx.parse([]string{"hello", "-123"})
	if ctx.lastError != nil {
		t.Error(ctx.lastError)
	}
}

func TestContext_parseOnce(t *testing.T) {
}
