package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	pid := os.Getpid()
	log.Println(pid)
	CalculateFd(fmt.Sprintf("%v", pid))
}

func CalculateFd(pid string) (fdNum uint, err error) {
	path := "/proc/" + pid + "/fd/"

	file, err := os.Open(path)
	if err != nil || nil == file {
		log.Println(err)
		return
	}
	defer file.Close()

	files, err := file.Readdirnames(0)

	if err != nil {
		return
	}

	fdNum = uint(len(files))

	return
}
