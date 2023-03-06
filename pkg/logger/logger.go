package logger

import (
	"github.com/lackone/go-ws/pkg/utils"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
)

type Logger struct {
	*zap.Logger
	folder string
	file   string
}

func NewLogger(folder, file string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) *Logger {
	if !utils.Exists(folder) {
		os.MkdirAll(folder, os.ModePerm)
	}

	logWriter := getLogWriter(filepath.Join(folder, file), maxSize, maxBackups, maxAge, compress)
	logEncoder := getLogEncoder()
	stackTraceLevel := getStackTraceLevel()

	core := zapcore.NewCore(logEncoder, logWriter, level)

	log := zap.New(core, zap.AddCaller(), zap.AddStacktrace(stackTraceLevel))
	zap.ReplaceGlobals(log)

	return &Logger{
		log,
		folder,
		file,
	}
}

func getStackTraceLevel() zap.LevelEnablerFunc {
	return func(level zapcore.Level) bool {
		return level >= zapcore.DPanicLevel
	}
}

func getLogWriter(filename string, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.WriteSyncer {
	lumberjack := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackups,
		MaxAge:     maxAge,
		Compress:   compress,
	}
	return zapcore.AddSync(lumberjack)
}

func getLogEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
