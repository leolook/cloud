package demo2

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	//"runtime"
	"fmt"
	"runtime"
	"strings"
)

type Test interface {
	Info(str ...interface{})
}

func CallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	//enc.AppendString(strings.Join([]string{caller.TrimmedPath()}, ":"))
	//enc.AppendString(strings.Join(runtime.Caller(0)))

	//enc.AppendString(strings.Join([]string{caller.TrimmedPath(), runtime.FuncForPC(caller.PC).Name()}, ":"))

	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		return
	}
	str := fmt.Sprintf("%v %s %d", pc, file, line)
	enc.AppendString(strings.Join([]string{str}, ":"))
}

func Info(str ...interface{}) {

	loggerConfig := zap.NewDevelopmentConfig()

	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {

		_, file, line, ok := runtime.Caller(0)
		if !ok {
			return
		}
		caller.File, caller.Line = file, line
		enc.AppendString(strings.Join([]string{caller.TrimmedPath()}, ":"))
	}

	logger, _ := loggerConfig.Build()

	logger.Info("huge")
}

type T struct {
	*zap.Logger
}
