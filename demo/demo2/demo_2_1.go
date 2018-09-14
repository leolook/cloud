package main

import (
	"cloud/lib/log"
)

type Test struct {
	Name string `json:"name"`
}

func main() {

	log.Infof("test=%+v", "huge")

	te := &Test{
		Name: "te",
	}

	log.Debugf("test=%+v", *te)
}
