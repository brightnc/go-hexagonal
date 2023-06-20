package logger

import (
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func Init() {
	var err error
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	loggerConfig.EncoderConfig.StacktraceKey = ""
	zapLog, err = loggerConfig.Build(zap.AddCallerSkip(1))
	if err != nil {
		log.Fatalln("Got error while building zap logger config.")
		return
	}

}

func Info(msg string, fields ...zapcore.Field) {
	zapLog.Info(msg, fields...)
}

func Debug(msg string, fields ...zapcore.Field) {
	zapLog.Debug(msg, fields...)
}

func Error(msg interface{}, fields ...zapcore.Field) {
	switch v := msg.(type) {
	case error:
		zapLog.Error(v.Error(), fields...)
	case string:
		zapLog.Error(v, fields...)
	}
}
