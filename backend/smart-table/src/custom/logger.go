package custom

import (
	"os"
	"sync"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerInstance *zap.Logger
	once           sync.Once
)

func InitLogger(cfg *Config) *zap.Logger {
	once.Do(func() {
		var level zapcore.Level
		switch cfg.Logging.Level {
		case DebugLevel:
			level = zap.DebugLevel
		case InfoLevel:
			level = zap.InfoLevel
		case WarnLevel:
			level = zap.WarnLevel
		case ErrorLevel:
			level = zap.ErrorLevel
		case FatalLevel:
			level = zap.FatalLevel
		default:
			level = zap.InfoLevel
		}

		var zapEncoderConfig zapcore.EncoderConfig
		var encoder zapcore.Encoder

		switch cfg.App.Env {
		case DevelopmentEnv:
			zapEncoderConfig = zap.NewDevelopmentEncoderConfig()
		case ProductionEnv:
			zapEncoderConfig = zap.NewProductionEncoderConfig()
		default:
			zapEncoderConfig = zap.NewDevelopmentEncoderConfig()
		}

		switch cfg.Logging.Format {
		case JsonFormat:
			encoder = zapcore.NewJSONEncoder(zapEncoderConfig)
		case ConsoleFormat:
			encoder = zapcore.NewConsoleEncoder(zapEncoderConfig)
		default:
			encoder = zapcore.NewConsoleEncoder(zapEncoderConfig)
		}

		logFile := &lumberjack.Logger{
			Filename:   cfg.Logging.File,
			MaxSize:    cfg.Logging.MaxSize,
			MaxBackups: cfg.Logging.MaxBackups,
			MaxAge:     cfg.Logging.MaxAge,
			Compress:   cfg.Logging.Compress,
		}

		core := zapcore.NewTee(
			zapcore.NewCore(encoder, zapcore.AddSync(logFile), level),
			zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), level),
		)

		loggerInstance = zap.New(core, zap.AddCaller())
	})

	return loggerInstance
}

func GetLogger() *zap.Logger {
	if loggerInstance == nil {
		panic("The logger is not initialised. Call InitLogger() before GetLogger()")
	}

	return loggerInstance
}
