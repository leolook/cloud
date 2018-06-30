package main

import "fmt"
import "net/url"

func main() {
	//我们将解析这个 URL 示例，它包含了一个 scheme，认证信息，主机名，端口，路径，查询参数和片段。
	s := "nl://gym/gym_info?gym_id=10334"
	//解析这个 URL 并确保解析没有出错。
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	//直接访问 scheme。
	fmt.Println(u.Scheme)
	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["gym_id"][0])
}
