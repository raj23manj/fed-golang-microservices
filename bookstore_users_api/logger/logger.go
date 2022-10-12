package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// naming `Log` will be public, but naming `log` will be private
// 21:10, No logging system
var (
	log *zap.Logger
)

func init() {
	logConfig := zap.Config{
		// logging in std o/p
		OutputPaths: []string{"stdout"},
		Level:       zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding:    "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:      "ts",
			LevelKey:     "level",
			MessageKey:   "msg",
			EncodeLevel:  zapcore.LowercaseLevelEncoder,
			EncodeTime:   zapcore.ISO8601TimeEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
	}

	var err error
	if log, err = logConfig.Build(); err != nil {
		panic(err)
	}
}

// overriding the default logger methods
// func Field(key string, value interface{}) zap.Field {
// 	return zap.Any(key, value)
// }

// func GetLogger() {
// 	return log
// }

func Info(msg string, tags ...zap.Field) {
	log.Info(msg, tags...)
	log.Sync()
}

// adding fields, 26:00, No logging system
func Error(msg string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("errorMysql", err))
	log.Error(msg, tags...)
	log.Sync()
}
