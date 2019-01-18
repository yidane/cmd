package internal

import (
	"strconv"
	"testing"
)

func TestDefaultParse_ParseArg(t *testing.T) {
	args, err := new(DefaultParse).ParseArg("./mdImage -rf    	/tmp/Database/ /sa/ -i -s=y -y")
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range args {
		t.Log("[", v.Key(), "=", v.Value(), "]")
	}
}

func Test_readStringAsList(t *testing.T) {
	arg := `1 2  3   4   5    6	7		8			9
							10   
11
12 13	14`

	l := readStringAsList(arg)

	node := l.Front()
	i := 1
	for node != nil {
		j, _ := strconv.Atoi(node.Value.(string))
		if i != j {
			t.Fatal(i, j)
		}
		node = node.Next()
		i++
	}

	if i != 15 {
		t.Fatal(i)
	}
}
