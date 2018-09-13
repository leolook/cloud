package kubectl

import (
	. "cloud/lib/util"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Kubectl struct {
	Env       string `json:"env"`
	Base      string `json:"base"`
	Namespace string `json:"-"`
}

//logs gym
//pods gym

func (k *Kubectl) Run(arg0 string, arg1 ...string) {

	if arg0 == "" {
		fmt.Printf("empty to arg0,arg0=%+v \n", arg0)
		return
	}

	if k.Env == "online" {
		k.Namespace = "--namespace=nl"
	} else {
		k.Namespace = fmt.Sprintf("--namespace=nl-%s", k.Env)
	}

	cmd := ""
	switch arg0 {
	case "pods":
		{
			if len(arg1) > 0 {
				cmd = k.pods(arg1[0])
			} else {
				cmd = k.pods("")
			}

			fmt.Printf("%s", Shell(cmd))

		}
	case "logs":
		{
			if len(arg1) <= 0 {
				fmt.Printf("empty to arg1,arg1=%+v \n,example:logs gym", arg0)
				return
			}

			name := k.findName(arg1[0])
			if name == nil || len(name) <= 0 {
				fmt.Printf("not found data,%s", arg1[0])
				return
			}

			arg1[0] = name[0]
			cmd = k.logs(arg1...)

			fmt.Printf("%s", Shell(cmd))

		}
	case "env":
		{
			if len(arg1) <= 0 {
				fmt.Printf("empty to arg1,arg1=%+v \n,example:env test", arg0)
				return
			}
			k.updateConfig(arg1[0])
		}
	}
}

func (k *Kubectl) updateConfig(arg string) {

	k.Env = arg
	byt, err := json.Marshal(k)
	if err != nil {
		fmt.Printf("failed to marshal,k=%+v,err=%+v\n", k, err)
		return
	}

	err = ioutil.WriteFile("config.json", byt, os.ModePerm)
	if err != nil {
		fmt.Printf("failed to write file,err=%+v\n", err)
		return
	}

	fmt.Printf("success to update env\n")
}

func (k *Kubectl) findName(arg string) []string {

	data := Shell(k.pods(arg))
	if !strings.Contains(data, "\n") {
		return nil
	}

	name := make([]string, 0)
	for _, v := range strings.Split(data, "\n") {

		tmp := []rune(v)
		for j, t := range tmp {
			if t == 32 && strings.Index(v, arg) == 0 {
				name = append(name, string(tmp[0:j]))
				break
			}
		}

	}

	return name
}

func (k *Kubectl) pods(arg string) string {

	order := ""
	if arg == "" {
		order = fmt.Sprintf("get pods %s", k.Namespace)
	} else {
		order = fmt.Sprintf("get pods %s | grep '%s'", k.Namespace, arg)
	}

	order = fmt.Sprintf("%s %s", k.Base, order)

	return order
}

//logs gym | grep 'ERROR'
func (k *Kubectl) logs(arg ...string) string {

	order := fmt.Sprintf("%s %s %s", "logs", arg[0], k.Namespace)

	if len(arg) >= 2 {
		for i := 1; i < len(arg); i++ {
			order += arg[i] + " "
		}
	}

	order = k.Base + " " + order

	return order
}
