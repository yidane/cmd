package log

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestGetLogger(t *testing.T) {
	log := GetLogger()

	var l Log
	if runtime.GOOS == "windows" {
		l = windowsLog{}
	} else {
		l = linuxLog{}
	}

	if log != l {
		t.Fatal("log should be ", reflect.TypeOf(l).Name())
	}
}

func Test_Errorf(t *testing.T) {
	log := GetLogger()
	log.Errorf("Errorf")
	log.Errorf("Errorf at %s", time.Now())
}

func Test_Error(t *testing.T) {
	log := GetLogger()
	log.Error(fmt.Errorf("Error"))
}

func Test_Succeed(t *testing.T) {
	log := GetLogger()
	log.Succeed("Succeed")
	log.Succeed("Succeed at %s", time.Now())
}

func Test_Warn(t *testing.T) {
	log := GetLogger()
	log.Warn("Warn")
	log.Warn("Warn at %s", time.Now())
}

func Test_Println(t *testing.T) {
	log := GetLogger()
	log.Println("Println")
	log.Println("Println", "at", time.Now())
}

func Test_Print(t *testing.T) {
	log := GetLogger()
	log.Print("Println", "\n")
	log.Print("Println", " ", "at", " ", time.Now())
}
