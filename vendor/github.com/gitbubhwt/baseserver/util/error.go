package util

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func Error(code int64, message string) error {
	return errors.New(fmt.Sprintf("%d;%s", code, message))
}

func ParseError(err string) (code int64, message string) {
	if err == "" {
		return 200, ""
	}
	if strings.Index(err, ";") == -1 {
		return 200, ""
	}
	val := strings.Split(err, ";")
	code, _ = strconv.ParseInt(val[0], 10, 64)
	return code, val[1]
}
