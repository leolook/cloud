package main

import (
	"cloud/lib/log"
	"reflect"
	"strings"
)

type Test struct {
	Name string `json:"name"`
}

func main() {

	te := &Test{
		Name: "胡文涛",
	}

	val := []*Test{te}

	typ := reflect.TypeOf(val)

	v := reflect.ValueOf(val)

	elem := v.
	log.Info(elem)
	if strings.Contains(elem.String(), "*") {
		elem = elem.Elem()
	}

	log.Info(elem.NumField(), elem.String(), elem)
	log.Info(typ, elem)

}
