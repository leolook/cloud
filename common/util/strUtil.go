package util

import (
	"fmt"
	"strings"
)

func StrRemoveSpace(str string) string {
	// 去除空格
	str = strings.Replace(str, " ", "", -1)
	// 去除换行符
	str = strings.Replace(str, "\n", "", -1)
	return str
}

//数组转化成逗号拼接的字符串  例如: [1,2,3] to "1,2,3"
func ArrayToStr(arr []interface{}) string {
	return strings.Replace(strings.Trim(fmt.Sprint(arr), "[]"), " ", ",", -1)
}
