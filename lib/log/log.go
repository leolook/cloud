package log

import (
	"cloud/lib/goid"
	"fmt"
	"sync"
)

// Debug logs to the DEBUG log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Debug(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Debugf(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Debug(args...)
	}
}

// Debugln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Debugln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Debugw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Debug(args...)
	}
}

// Debugf logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Debugf(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Sugar().Debugw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Debugf(format, args...)
	}
}

// Info logs to the INFO log.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Info(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Infow(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Info(args...)
	}
}

// Infoln logs to the INFO log.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Infoln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Infow(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Info(args...)
	}
}

// Infof logs to the INFO log.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Infof(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Sugar().Infow(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Infof(format, args...)
	}
}

func InfoW(format string, keysAndValues ...interface{}) {
	g_logger.Sugar().Infow(format, keysAndValues...)
}

// Warning logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Warn(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Warnw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Warn(args...)
	}
}

func Warnw(format string, keysAndValues ...interface{}) {
	g_logger.Sugar().Warnw(format, keysAndValues...)
}

// Warningln logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Warnln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Warnw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Warn(args...)
	}
}

// Warningf logs to the WARNING and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Warnf(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Sugar().Warnw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Warnf(format, args...)
	}
}

// Error logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Error(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Error(args...)
	}
}

// Errorln logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Errorln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Error(args...)
	}
}

// Errorf logs to the ERROR, WARNING, and INFO logs.
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Errorf(format string, args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Sugar().Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Errorf(format, args...)
	}
}

func Errorw(format string, keysAndValues ...interface{}) {
	g_logger.Sugar().Errorw(format, keysAndValues...)
}

// Fatal logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Print; a newline is appended if missing.
func Fatal(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Errorw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Fatal(args...)
	}
}

// Fatalln logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Println; a newline is appended if missing.
func Fatalln(args ...interface{}) {
	gid := goid.Get()
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprint(args)
		g_logger.Sugar().Fatalw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Fatal(args...)
	}
}

// Fatalf logs to the FATAL, ERROR, WARNING, and INFO logs,
// including a stack trace of all running goroutines, then calls os.Exit(255).
// Arguments are handled in the manner of fmt.Printf; a newline is appended if missing.
func Fatalf(format string, args ...interface{}) {
	gid := goid.Get()
	fmt.Println(gid)
	if color, ok := GetColorByKey(gid); ok {
		msg := fmt.Sprintf(format, args...)
		g_logger.Sugar().Fatalw(msg, "method", color.Method(), "traceid", color.TraceID(), "uid", color.Uid(), "clientversion", color.ClientVerson(), "clienttype", color.ClentType())
	} else {
		g_logger.Sugar().Fatalf(format, args...)
	}
}

type colorI interface {
	String() string
	Uid() string
	ClientVerson() string
	Method() string
	TraceID() string
	ClentType() string
}

var g_color sync.Map

func SetColor(val colorI) bool {
	gid := goid.Get()
	_, loaded := g_color.LoadOrStore(gid, val)
	return !loaded
}

func GetColorByKey(key int64) (colorI, bool) {
	val, ok := g_color.Load(key)
	if !ok {
		return nil, ok
	} else {
		return val.(colorI), ok
	}
}

func GetColor() (colorI, bool) {
	gid := goid.Get()
	val, ok := g_color.Load(gid)
	return val.(colorI), ok
}

func DelColor() {
	gid := goid.Get()
	g_color.Delete(gid)
}
