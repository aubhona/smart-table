package logging

import (
	"os"
	"sync"

	"github.com/smart-table/src/config"

	"github.com/natefinch/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	loggerInstance *zap.Logger
	once           sync.Once
)

func InitLogger(cfg *config.Config) *zap.Logger {
	once.Do(func() {
		var level zapcore.Level

		switch cfg.Logging.Level {
		case config.DebugLevel:
			level = zap.DebugLevel
		case config.InfoLevel:
			level = zap.InfoLevel
		case config.WarnLevel:
			level = zap.WarnLevel
		case config.ErrorLevel:
			level = zap.ErrorLevel
		case config.FatalLevel:
			level = zap.FatalLevel
		default:
			level = zap.InfoLevel
		}

		var zapEncoderConfig zapcore.EncoderConfig

		var encoder zapcore.Encoder

		switch cfg.App.Env {
		case config.DevelopmentEnv:
			zapEncoderConfig = zap.NewDevelopmentEncoderConfig()
		case config.ProductionEnv:
			zapEncoderConfig = zap.NewProductionEncoderConfig()
		default:
			zapEncoderConfig = zap.NewDevelopmentEncoderConfig()
		}

		switch cfg.Logging.Format {
		case config.JSONFormat:
			encoder = zapcore.NewJSONEncoder(zapEncoderConfig)
		case config.ConsoleFormat:
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
		panic("The logger is not initialized. Call InitLogger() before GetLogger()")
	}

	return loggerInstance
}
