package server

import (
	"cloud/lib/log"
	"reflect"
)

type Server struct {
	Name    string
	Addr    string
	Service interface{}
}

func (s *Server) Run() {

	t := reflect.TypeOf(s.Service)

	for i := 0; i < t.NumMethod(); i++ {

		me := t.Method(i)

		meType := me.Type

		if meType.NumIn() > 1 {
			log.Fatalf("input parameter is not in 0~1,len=%d", meType.NumIn())
			break
		}

		if meType.NumOut() < 1 || meType.NumOut() > 2 {
			log.Fatalf("output parameter is not in 1~2,len=%d", meType.NumOut())
			break
		}

		if meType.NumOut() == 1 && meType.Out(0).Name() != "error" {
			log.Fatalf("output parameter is not error when only one parameter,len=%+v,name=%s", meType.NumOut(), meType.Out(0).Name())
			break
		}

	}

}
