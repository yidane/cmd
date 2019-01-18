package internal

import (
	"testing"
)

func TestDefaultParse_ParseArg(t *testing.T) {
	args, err := new(DefaultParse).ParseArg("./mdImage -rf /tmp/Database/")
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range args {
		t.Log(v.Key(), "=", v.Value())
	}
	//t.Log("nothing implement")
}
