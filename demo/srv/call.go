package srv

import "fmt"

type Call interface {
	BaseSrv
	Call(data string)
}

type BaseSrv interface {
	Do(data string)
}

type Base struct {
}

func (d Base) Do(data string) {
   fmt.Println(data)
}

type Mobile struct {
	Base
}

func (m Mobile) Call(data string) {
	fmt.Println(fmt.Sprintf("Call mobile,%v", data))
}

type Phone struct {
	Base
}

func (p Phone) Call(data string) {
	fmt.Println(fmt.Sprintf("Call phone,%v", data))
	p.Do(data)
}
