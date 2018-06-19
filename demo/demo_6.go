package main

import "fmt"

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
    Test_1().Print("hwt")
}
