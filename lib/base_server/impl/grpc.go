package impl

import (
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"reflect"
)

type GrpcImpl struct {
	Name string
	Addr string      //host + port
	Srv  interface{} //服务对象 实现grpc方法的对象
}

func (impl *GrpcImpl) StartServer() error {
	if impl.Addr == "" || impl.Srv == nil {
		return errors.New(fmt.Sprintf("Invalid argumet,[addr=%v] [srv=%v]", impl.Addr, impl.Srv))
	}

	listen, err := net.Listen("tcp", impl.Addr)
	if err != nil {
		return err
	}

	t := reflect.TypeOf(impl.Srv)
	v := reflect.ValueOf(impl.Srv)

	isReg := false
	methodRegName := "RegServer"

	for i := 0; i < t.NumMethod(); i++ {

		methodName := t.Method(i).Name
		method := v.Method(i)
		mtype := method.Type()

		if methodName == methodRegName && mtype.NumIn() == 1 {
			isReg = true
			continue
		}

		if mtype.NumIn() != 2 {
			return errors.New(fmt.Sprintf("%s input params size is wrong,[size=%v]", methodName, mtype.NumIn()))
		}

		if mtype.In(0).String() != "context.Context" {
			return errors.New(fmt.Sprintf("%s first input param type is wrong,[type=%v]", methodName, mtype.In(0)))
		}

		if mtype.Out(1).String() != "error" {
			return errors.New(fmt.Sprintf("%s second out param type is wrong,[type=%v]", methodName, mtype.Out(1)))
		}
	}

	go func() {
		server := grpc.NewServer()
		if isReg {
			arg := make([]reflect.Value, 1)
			arg[0] = reflect.ValueOf(server)
			v.MethodByName(methodRegName).Call(arg)
		}
		server.Serve(listen)
	}()

	return nil
}
