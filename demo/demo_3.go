package main

import (
	"regexp"
	"log"
)

func main(){
	match, _ := regexp.MatchString("/cloud/admin/([a-z]+)", "/cloud/admin/admin/aaaa")
	log.Println(match)
}
