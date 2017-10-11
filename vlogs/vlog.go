package vlogs

import (
	"go.uber.org/zap"
	"time"
	"gopkg.in/natefinch/lumberjack.v2"
	"go.uber.org/zap/zapcore"
	"beego"
)

var VesyncLog = newLogger(beego.AppConfig.String("serverName"),beego.AppConfig.String("logFilePath"),time.Duration(time.Hour*24))

func newLogger(serverName string, logFilePath string, rotationTime time.Duration) *zap.Logger{
	l := &lumberjack.Logger{Filename: logFilePath, MaxAge: 30, MaxSize: 1024 * 10}
	w := zapcore.AddSync(l)

	var logLevel zapcore.Level
	switch beego.AppConfig.String("logLevel") {
	case "debug":
		logLevel=zap.DebugLevel
	case "info":
		logLevel=zap.InfoLevel
	case "warn":
		logLevel=zap.WarnLevel
	case "error":
		logLevel=zap.ErrorLevel
	default:
		logLevel=zap.DebugLevel
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     utcTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}), w, logLevel)

	vLog := zap.New(core).Named(serverName)

	go func() {
		defer VesyncLog.Sync()
		t := time.NewTicker(rotationTime)
		for {
			select {
			case <-t.C:
				l.Rotate()
			}
		}
	}()
	return vLog
}

func utcTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.00000"))
}

