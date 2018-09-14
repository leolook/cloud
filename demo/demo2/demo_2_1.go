package main

import (
	"cloud/lib/log"
	"cloud/lib/util"
)

type Test struct {
	Name string `json:"name"`
}

func main() {

	te := &Test{
		Name: "te",
	}

	data := util.Shell("ls")

	log.Infof("%+v,%+v", te, data)

}
