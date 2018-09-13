package util

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Shell(s string) string {

	fmt.Printf("%s \n", s)

	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return out.String()
}
