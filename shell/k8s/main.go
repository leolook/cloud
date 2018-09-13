package main

import (
	"cloud/shell/k8s/kubectl"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var addr string

func init() {
	flag.StringVar(&addr, "addr", "/Users/huwentao/me/k8s/k8s_env/config.json", "config")
}

func main() {

	flag.Parse()

	fmt.Printf("config path:%s\n", addr)

	input := fetchInput()
	if input == nil || len(input) <= 0 {
		return
	}

	k := fetchConfig()
	if k == nil {
		return
	}

	if len(input) >= 3 {
		k.Run(input[1], input[2:]...)
	} else {
		k.Run(input[1])
	}
}

func fetchInput() []string {

	input := make([]string, 0)

	for i := range os.Args {
		input = append(input, os.Args[i])
	}

	if len(input) <= 1 {
		fmt.Printf("logs \npods \n")
		return nil
	}

	return input
}

func fetchConfig() *kubectl.Kubectl {

	byt, err := ioutil.ReadFile(addr)
	if err != nil {
		fmt.Printf("failed to read file,err=%+v\n", err)
		return nil
	}

	var kub kubectl.Kubectl
	err = json.Unmarshal(byt, &kub)
	if err != nil {
		fmt.Printf("failed to unmarshal,err=%+v", err)
		return nil
	}

	return &kub

}
