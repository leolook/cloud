package log

import (
	"flag"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

const (
	DEBUG = iota
	INFO
	WARN
	ERROR
	FATAL
)

// MaxSize is the maximum size of a log file in bytes.
var MaxSize int = 1024 * 1024 * 1800
var MaxFileNum int = 10

// If non-empty, overrides the choice of directory in which to write logs.
// See createLogDirs for the full list of possible destinations.
var logDir string

var g_logger *zap.Logger

var LogFileName string = ""

var LogLevel string = "debug"

var LogEnCode string = "json"

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func createLogDirs() error {

	exist, err := PathExists(logDir)
	if err != nil {
		return err
	}

	if !exist {
		// 创建文件夹
		return os.MkdirAll(logDir, os.ModePerm)
	}
	return nil
}

func init_xlog() error {

	var flag bool = true
	err := createLogDirs()
	if err != nil {
		flag = false
	}

	logcfg := NewDefaultLoggerConfig()
	g_logger = logcfg.NewLogger(flag)

	return nil
}

func fin_xlog() error {
	g_logger.Sugar().Sync()
	return nil
}

func init() {

	flag.IntVar(&MaxSize, "max_log_size", 10*1024*1024, "max log file size")
	flag.IntVar(&MaxFileNum, "max_log_num", 10, "max log file num")

	flag.StringVar(&LogLevel, "log_level", "debug", "defalut debug")
	flag.StringVar(&LogFileName, "log_file", "server.log", "log file name")

	flag.StringVar(&logDir, "log_dir", "/data/logs/", "log file dir")

	flag.StringVar(&LogEnCode, "log_encode", "json", "log defalut encode")

	flag.Parse()

	init_xlog()
}

//LoggerConfig config of logger
type LoggerConfig struct {
	Level      string //debug  info  warn  error
	Encoding   string //json or console
	CallFull   bool   //whether full call path or short path, default is short
	Filename   string //log file name
	MaxSize    int    //max size of log.(MB)
	MaxAge     int    //time to keep, (day)
	MaxBackups int    //max file numbers
	LocalTime  bool   //(default UTC)
	Compress   bool   //default false
}

func convertLogLevel(levelStr string) (level zapcore.Level) {
	switch levelStr {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	}
	return
}

//NewDefaultLoggerConfig create a default config
func NewDefaultLoggerConfig() *LoggerConfig {
	fname := logDir + LogFileName
	return &LoggerConfig{
		Level:      LogLevel,
		Filename:   fname,
		MaxSize:    MaxSize,
		MaxAge:     1,
		MaxBackups: MaxFileNum,
		Encoding:   LogEnCode,
	}
}

func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.999999999"))
}

func (lconf *LoggerConfig) NewLogger(flag bool) *zap.Logger {

	if lconf.Filename == "" {
		logger, _ := zap.NewProduction(zap.AddCallerSkip(1))
		return logger
	}

	config := zap.NewProductionConfig()
	//config := zap.NewDevelopmentConfig()
	config.EncoderConfig.EncodeTime = timeEncoder

	var encoder zapcore.Encoder
	if lconf.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(config.EncoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(config.EncoderConfig)
	}

	//zapWriter := zapcore.
	zapWriter1 := zapcore.AddSync(&lumberjack.Logger{
		Filename:   lconf.Filename,
		MaxSize:    lconf.MaxSize,
		MaxAge:     lconf.MaxAge,
		MaxBackups: lconf.MaxBackups,
		LocalTime:  lconf.LocalTime,
	})

	var zapWriter zapcore.WriteSyncer
	if flag {
		zapWriter = zapcore.NewMultiWriteSyncer(zapWriter1, os.Stdout)
	} else {
		zapWriter = zapcore.AddSync(os.Stdout)
	}

	newCore := zapcore.NewCore(encoder, zapWriter, zap.NewAtomicLevelAt(convertLogLevel(lconf.Level)))
	opts := []zap.Option{}
	opts = append(opts, zap.AddCaller(), zap.AddCallerSkip(1))
	logger := zap.New(newCore, opts...)

	return logger
}

func GetLevel() int {
	switch LogLevel {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "error":
		return ERROR
	case "warn":
		return WARN
	}
	return 0
}
