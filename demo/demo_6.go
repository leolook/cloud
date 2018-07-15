package main

import "fmt"

type Render interface {
	Eat(data string)
}

type People interface {
	Instance(name, data string) Render
}

type (
	Child struct {
		Name string
		Age  int64
	}

	Old struct {
		Name string
	}
)

func (r Child) Instance(name, data string) Render {
	return Child{
		Name: "j",
		Age:  1,
	}
}

func (r Child) Eat(data string) {
	fmt.Println(r.Name, "==>", r.Age)
}

func (r Old) Instance(name, data string) Render {
	return Old{
		Name: "j",
	}
}

func (r Old) Eat(data string) {
	fmt.Println(r.Name)
}

type Do struct {
	Render People
}

func (d Do) Do(name, data string) {
	r := d.Render.Instance(name, data)
	r.Eat(data)
}
func main() {
	do := Do{
		Render: Old{},
	}

	do.Do("nihao", "huhu")
}
