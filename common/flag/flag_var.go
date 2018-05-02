package flag

import "flag"

var ConfPath string
var FilePath string

func init() {
	flag.StringVar(&ConfPath, "conf", "config.yml", "config path")
	flag.StringVar(&FilePath, "file", "file", "file path")
	flag.Parse()
}
