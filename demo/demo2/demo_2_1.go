package main

import (
	"cloud/lib/log"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
)

type Context map[string]interface{}

func (c *Context) get(key string) interface{} {
	mp := map[string]interface{}(*c)
	return mp[key]
}

func (c *Context) put(key string, val interface{}) {
	mp := make(map[string]interface{})
	mp[key] = val
	*c = mp
}


func main() {

	eng := gin.Default()
	r := eng.Group("demo")

	s := &Controller{}
	typ := reflect.TypeOf(s)
	val := reflect.ValueOf(s)

	for i := 0; i < typ.NumMethod(); i++ {

		mt := typ.Method(i)
		mv := val.Method(i)

		if mv.Type().NumOut() <= 0 {
			log.Fatalf("%s method have not out put parameter", mt.Name)
			break
		}

		if mv.Type().NumOut() > 1 {
			log.Fatalf("%s method can only have one out put parameter", mt.Name)
			break
		}

		if mv.Type().NumIn() > 2 {
			log.Fatalf("%s method can only less than two in parameter", mt.Name)
			break
		}

		for i := 0; i < mv.Type().NumIn(); i++ {
			name := mv.Type().In(i).String()
			if !isStruct(name) {
				log.Fatalf("%s method input parameter can only is struct,name=%s", mt.Name, name)
				break
			}
		}

		if mv.Type().NumIn() == 2 {
			isCtx := false
			for i := 0; i < mv.Type().NumIn(); i++ {
				name := mv.Type().In(i).String()
				if strings.Contains(name, ".Context") {
					isCtx = true
					break
				}
			}
			if !isCtx {
				log.Fatalf("%s method has no context when two input parameter", mt.Name)
				break
			}
		}

		r.POST(mt.Name, func(ctx *gin.Context) {

			//接口没有输入
			var rsp []reflect.Value
			if mv.Type().NumIn() == 0 {
				rsp = mv.Call(nil)
			} else {
				//var req []reflect.Value
				ct := ctx.GetHeader("Content-Type")
				log.Info(ct)
				if strings.Contains(ct, "multipart/form-data") {

				} else if strings.Contains(ct, "application/json") {

				} else if strings.Contains(ct, "application/x-www-form-urlencoded") {

				}
			}

			//接口返回数据
			if rsp == nil || len(rsp) <= 0 {
				ctx.JSON(http.StatusInternalServerError, "server error")
				ctx.Abort()
				return
			}

			out := mv.Type().Out(0)

			switch out.Name() {
			case "int", "int8", "int32", "int64", "uint", "uint8", "uint32", "uint64", "float32", "float64", "string", "bool":
				{
					ctx.String(http.StatusOK, "%v", rsp[0].Interface())
				}
			default:
				{
					by, err := json.Marshal(rsp[0].Interface())
					if err != nil {
						log.Errorf("failed to unmarshal,err=%v", err)
						ctx.JSON(http.StatusInternalServerError, "server error")
						ctx.Abort()
						return
					}
					ctx.Status(http.StatusOK)
					ctx.Writer.Write(by)
				}

			}

		})

	}

	eng.Run("0.0.0.0:1233")
}

func isStruct(name string) bool {
	switch name {
	case "int", "int8", "int32", "int64", "uint", "uint8", "uint32",
		"uint64", "float32", "float64", "string", "bool", "interface{}":
		{
			return false
		}
	default:
		{
			return true
		}
	}
}

type Controller struct {
}

type RegReq struct {
	ID int64 `json:"id"`
}

//@get
func (c *Controller) Login() string {
	return "login"
}

func (c *Controller) Register(req *RegReq) string {

	return "register"
}

func (c *Controller) Register1(req *RegReq, ctx *Context) map[string]interface{} {

	mp := make(map[string]interface{})
	mp["test"] = 1
	return mp
}
