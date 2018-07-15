package main

import (
	"fmt"
	"reflect"
)

func main() {
	tonydon := &User{Name: "TangXiaodong", Age: 100, Id: "0000123"}
	object := reflect.ValueOf(tonydon)
	myref := object.Elem()
	typeOfType := myref.Type()
	for i := 0; i < myref.NumField(); i++ {
		field := myref.Field(i)
		fmt.Printf("%d. %s %s = %v \n", i, typeOfType.Field(i).Name, field.Type(), field.Interface())
	}
	tonydon.SayHello()
	v := object.MethodByName("SayHello")
	v.Call([]reflect.Value{})
}

type User struct {
	Name string
	Age  int
	Id   string
	Arr  []string
}

func (u *User) SayHello() {
	fmt.Println("I'm " + u.Name + ", Id is " + u.Id + ". Nice to meet you! ")
}
