package dispatch

import (
	"cloud/demo/srv"
	"fmt"
)

type DispatchSrv interface {
	Do(id int64, content string)
}

var (
	do map[int64]srv.Call
)

func init() {
	do = make(map[int64]srv.Call)
	do[0] = srv.Phone{}
	do[1] = srv.Mobile{}
}

//调度
type Dispatch struct {
	//map[int64]
}

//执行调度
func (t Dispatch) Do(id int64, content string) {
	if v, ok := do[id]; ok {
		v.Call(content)
	} else {
		fmt.Println("Not found dispatch,id:", id)
	}
}
