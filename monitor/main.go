package main

import (
	"cloud/lib/http/server"
	"cloud/lib/log"
)

func main() {

	s := server.Server{
		Name:    "monitor",
		Addr:    "0.0.0.0:10322",
		Service: &Controller{},
	}

	s.Run()
}

type Controller struct{}

type RegReq struct {
	ID int64 `json:"id"`
}

//@get
func (c *Controller) Login() string {
	return "login"
}

func (c *Controller) Register(ctx *server.Context) string {

	log.Info(ctx)
	return ctx.Get("name")
}

func (c *Controller) Register1(req *RegReq, ctx *server.Context) map[string]interface{} {

	mp := make(map[string]interface{})
	mp["test"] = ctx.Get("name")
	return mp
}
