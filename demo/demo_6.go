package main

import (
	"fmt"
	"log"
	"strings"
)

type Srv interface {
	Print(data string)
}

type Func func(data string)

func (f Func) Print(data string) {
	f(data)
}

func Test_1() Srv {
	return Func(func(data string) {
		fmt.Println("Test_1:", data)
	})
}

func main() {
	//Test_1().Print("hwt")

	str := ""
	for i := 0; i < 3; i++ {
		str = strings.Join([]string{str}, ",")
	}
	log.Println(str)

	var st *int = nil

	var h int = *st

	log.Println(h)
}
