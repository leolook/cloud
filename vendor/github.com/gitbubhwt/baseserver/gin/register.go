package gin

import (
	"fmt"
	log "github.com/alecthomas/log4go"
	"github.com/gin-gonic/gin"
	pb "github.com/gitbubhwt/baseserver/protocol"
	"github.com/gitbubhwt/baseserver/util"
	"net/http"
	"reflect"
	"strings"
)

type CommonReqFunc func(ctx *gin.Context) (interface{}, error)

type BaseServer struct {
	Server        interface{}
	Group         *gin.RouterGroup
	Middle        gin.HandlerFunc
	CommonReqFunc CommonReqFunc
}

func (b *BaseServer) Register() {

	if b.Group == nil || b.Server == nil {
		log.Info(fmt.Sprintf("Invalid input params,[server=%v] [group=%v]", b.Server, b.Group))
		return
	}

	server, g := b.Server, b.Group

	if b.Middle != nil {
		g.Use(b.Middle)
	}

	t := reflect.TypeOf(server)
	v := reflect.ValueOf(server)
	for i := 0; i < t.NumMethod(); i++ {

		methodName := t.Method(i).Name
		tmp := []rune(methodName)
		methodName = strings.ToLower(string(tmp[0])) + string(tmp[1:])
		method := v.Method(i)

		//输入、输出参数校验
		methodType := method.Type()
		//输出参数
		if methodType.NumOut() != 2 {
			log.Error(fmt.Sprintf("Method num out is wrong,[numOut=%v]", methodType.NumOut()))
			continue
		}

		if methodType.Out(1).String() != "error" {
			log.Error(fmt.Sprintf("Method out params,second out param is not error,[out=%v]", methodType.Out(1)))
			continue
		}

		//输入参数
		if methodType.NumIn() <= 0 || methodType.NumIn() >= 3 {
			log.Error(fmt.Sprintf("Method num in is wrong,[numIn=%v]", methodType.NumIn()))
			continue
		}

		var commonType, reqType reflect.Type

		if methodType.NumIn() == 1 {
			reqType = methodType.In(0)
		} else if methodType.NumIn() == 2 {
			commonType = methodType.In(0)
			reqType = methodType.In(1)
		}

		g.POST(methodName, func(ctx *gin.Context) {

			//数据绑定
			req := reflect.New(reqType)
			err := ctx.Bind(req.Interface())
			if err != nil {
				log.Error(fmt.Sprintf("Bind req fail,[err=%v]", req))
				ctx.JSON(http.StatusOK, util.FailRsp(pb.ERR_INPUT_WRONG, "Input data wrong"))
				return
			}

			reqElem := req.Elem()
			//参数校验 是否空
			rsp := checkField(reqElem)
			if rsp != "" {
				ctx.JSON(http.StatusOK, util.FailRsp(pb.ERR_IS_EMPTY, rsp))
				return
			}

			//执行调用
			var args []reflect.Value
			if commonType != nil && b.CommonReqFunc != nil {
				//执行公共模块
				common, err := b.CommonReqFunc(ctx)
				if err != nil {
					code, message := util.ParseError(fmt.Sprintf("%v", err))
					ctx.JSON(http.StatusOK, util.FailRsp(code, message))
					return
				}
				if common == nil {
					panic("common is nil")
				}
				args = append(args, reflect.ValueOf(common))
			}
			args = append(args, reqElem)

			doCall(ctx, method, args)
		})
	}
}

//参数校验 是否空
func checkField(elem reflect.Value) string {
	var rsp string
	for i := 0; i < elem.Type().NumField(); i++ {

		field := elem.Type().Field(i)
		empty := field.Tag.Get("empty")
		if empty != "no" {
			continue
		}

		pass := true
		value := elem.Field(i)
		//log.Info(fmt.Sprintf("value=%v", value.Interface()))

		switch value.Interface().(type) {
		case uint, uint32, uint64, int, int32, int64:
			{
				if value.Int() <= 0 {
					pass = false
				}
				break
			}
		case float32, float64:
			{
				if value.Float() <= 0 {
					pass = false
				}
				break
			}
		case string:
			{
				if value.String() == "" {
					pass = false
				}
			}
		case []string, []float64, []float32, []int64, []int, []int32, []uint32, []uint16, []uint64, []interface{}:
			{

				if value.Cap() == 0 {
					pass = false
				}
				break
			}
		}
		if pass {
			continue
		}
		rsp = fmt.Sprintf("%s is empty", field.Tag.Get("json"))
		break
	}
	return rsp
}

//执行调用
func doCall(ctx *gin.Context, method reflect.Value, args []reflect.Value) {
	val := method.Call(args)

	//log.Error(val[0].Interface(), val[1].Interface())
	//调用后输出
	if len(val) <= 0 || len(val) >= 3 {
		log.Error(fmt.Sprintf("Method call return value is wrong,[val=%v]", val))
		ctx.JSON(http.StatusOK, util.FailRsp(pb.ERR_SERVER_ERR, "server error"))
		return
	}

	if val[1].Interface() != nil {
		code, message := util.ParseError(fmt.Sprintf("%v", val[1].Interface()))
		ctx.JSON(http.StatusOK, util.FailRsp(code, message))
		return
	}

	if val[0].Interface() == nil {
		ctx.JSON(http.StatusOK, util.FailRsp(pb.ERR_NOT_FOUND, "not found data"))
		return
	}

	ctx.JSON(http.StatusOK, util.SucRsp(200, val[0].Interface()))
}
