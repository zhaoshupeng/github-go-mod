package helper

import (
	"fmt"

	"github.com/zhaoshupeng/golib"
	"github.com/zhaoshupeng/golib/ty"
)

type Help struct {
	Name string
}

func Echo() {
	fmt.Println("echo help")
	golib.LibEcho()
	fmt.Println(ty.Dir)

}
