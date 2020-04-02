package main

import (
	"fmt"

	"github.com/ActiveState/vt10x"
)

func main() {
	strip := vt10x.NewStrip()
	res, _ := strip.Strip([]byte("\033[?25hhello\033[97m\033[38X\033[1;43H world"))
	fmt.Println(string(res))
}
