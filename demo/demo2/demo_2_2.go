package main

import (
	"fmt"
)

type Base struct {
	Id int64
}

type People struct {
	Base
	Name string
}

func (p *People) Eat(str string) {
	fmt.Println(p.Name, str)
}

func main() {

	p := &People{
		Name: "1",
	}

	p.Eat("dd")

}
