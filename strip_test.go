package vt10x_test

import (
	"testing"

	"github.com/ActiveState/vt10x"
)

func TestStrip(t *testing.T) {

	strip := vt10x.NewStrip()
	res, _ := strip.Strip([]byte("\033[?25hhello\033[97m\033[38X\033[1;43H world"))
	if string(res) != "hello world" {
		t.Errorf("expected res to equal 'hello world', but was %s\n", string(res))
	}
}
