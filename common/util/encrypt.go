package util

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
)

func Sha1Md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
	fmt.Println(md5str)
	sha := sha1.Sum([]byte(md5str))
	return fmt.Sprintf("%x", sha)
}
