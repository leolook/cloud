package server

import (
	"cloud/lib/log"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/json"
	"net/http"
	"reflect"
	"strings"
)

type Server struct {
	Name    string
	Addr    string
	Service interface{}
	route   *gin.RouterGroup
}

func (s *Server) Run() {

	eng := gin.New()
	s.route = eng.Group(s.Name)

	t := reflect.TypeOf(s.Service)
	v := reflect.ValueOf(s.Service)
	for i := 0; i < t.NumMethod(); i++ {

		mv, mt := v.Method(i), t.Method(i)

		//参数检查
		isPass := s.parameterCheck(mt, mv)
		if !isPass {
			break
		}

		//接口注入
		s.interfaceInject(mv, mt)
	}

	eng.Run(s.Addr)
}

//参数检查
func (s *Server) parameterCheck(mt reflect.Method, mv reflect.Value) bool {

	if mv.Type().NumOut() <= 0 {
		log.Fatalf("%s method have not out put parameter", mt.Name)
		return false
	}

	if mv.Type().NumOut() > 1 {
		log.Fatalf("%s method can only have one out put parameter", mt.Name)
		return false
	}

	if mv.Type().NumIn() > 2 {
		log.Fatalf("%s method can only less than two in parameter", mt.Name)
		return false
	}

	for i := 0; i < mv.Type().NumIn(); i++ {
		name := mv.Type().In(i).String()
		if !s.isStruct(name) {
			log.Fatalf("%s method input parameter can only is struct,name=%s", mt.Name, name)
			return false
		}
	}

	if mv.Type().NumIn() == 2 && !s.isCtx(mv) {
		log.Fatalf("%s method has no context when two input parameter", mt.Name)
		return false
	}

	return true
}

//是否是结构体
func (s *Server) isStruct(name string) bool {

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

//当两个参数时，是否包含上下文参数
func (s *Server) isCtx(mv reflect.Value) bool {

	for i := 0; i < mv.Type().NumIn(); i++ {
		name := mv.Type().In(i).String()
		if strings.Contains(name, ".Context") {
			return true
		}
	}

	return false
}

//接口注入
func (s *Server) interfaceInject(mv reflect.Value, m reflect.Method) {

	s.route.POST(s.formatName(m.Name), func(ctx *gin.Context) {

		//接口没有输入
		var rsp []reflect.Value
		if mv.Type().NumIn() == 0 {
			rsp = mv.Call(nil)
		} else {

			inputSize := mv.Type().NumIn()
			req := make([]reflect.Value, inputSize)

			for i := 0; i < inputSize; i++ {
				name := mv.Type().In(i).String()
				if strings.Contains(name, ".Context") {
					context := new(Context)
					for k, v := range ctx.Request.Header {

						ru := []rune(k)
						if ru[0] >= 65 && ru[0] <= 90 { //大写字母
							ru[0] += 32
						}

						k = string(ru)

						context.Put(k, v[0])
					}
					req[i] = reflect.ValueOf(context)

				} else {
					param := reflect.New(mv.Type().In(i)).Elem()

					ct := ctx.GetHeader("Content-Type")
					if strings.Contains(ct, "application/json") {

						err := ctx.BindJSON(&param)
						if err != nil {
							log.Errorf("failed to bind json,err=%v", err)
							ctx.JSON(http.StatusInternalServerError, "server error")
							ctx.Abort()
							return
						}

					} else if strings.Contains(ct, "application/x-www-form-urlencoded") {

						log.Infof("%+v", param)

						//for i := 0; i < param.Type().NumField(); i++ {
						//	//log.Info(param.Field(i).String())
						//
						//	log.Info(param.Type().Field(i).Tag)
						//}

					}

					req[i] = param
				}
			}

			rsp = mv.Call(req)
		}

		//接口输出
		s.makeRsp(mv, rsp, ctx)
	})
}

func (s *Server) makeRsp(mv reflect.Value, rsp []reflect.Value, ctx *gin.Context) {
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
}

func (s *Server) formatName(str string) string {
	ru := []rune(str)
	tmp := make([]rune, 0)
	for i, v := range ru {
		if v >= 65 && v <= 90 { //大写字母
			v += 32
			if i > 0 && i < len(ru)-1 {
				tmp = append(tmp, 47)
			}
		}
		tmp = append(tmp, v)
	}
	return string(tmp)
}
