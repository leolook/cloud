package main

import (
	"fmt"
	"regexp"
)

func main() {
	reg := regexp.MustCompile("/cloud")
	//reg := regexp.MustCompile("/cloud/([a-z]+)ch")

	num := reg.FindStringIndex("/cloud/shelvess/")
	suc := reg.MatchString("/test/cloud/shelvesch")

	res := reg.ReplaceAllString("/cloud/shelvess/", "")
	fmt.Println(suc, num, res)

}
