package config

import (
	"context"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logDir = "/var/log/thc"

type loggerCtxKey struct{}

var loggerKey = loggerCtxKey{}

var (
	infoFileName  = "info.log"
	debugFileName = "debug.log"
	errorFileName = "error.log"
)

func fileRotation(
	filePath string,
) zapcore.WriteSyncer {

	return zapcore.AddSync(
		&lumberjack.Logger{
			Filename:   filePath,
			MaxSize:    10,
			MaxBackups: 5,
			MaxAge:     7,
		},
	)
}

func createEncoder() zapcore.Encoder {

	cfg :=
		zap.NewProductionEncoderConfig()

	cfg.EncodeTime =
		zapcore.ISO8601TimeEncoder

	cfg.TimeKey =
		"timestamp"

	cfg.EncodeLevel =
		zapcore.CapitalLevelEncoder

	cfg.EncodeCaller =
		zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(cfg)
}

func createWriteSyncer(
	filePath string,
) zapcore.WriteSyncer {

	if err :=
		os.MkdirAll(
			filepath.Dir(filePath),
			0755,
		); err != nil {

		panic(err)
	}

	return fileRotation(
		filePath,
	)
}

func createCore(
	level zapcore.LevelEnabler,
	file string,
) zapcore.Core {

	return zapcore.NewCore(
		createEncoder(),

		createWriteSyncer(
			filepath.Join(
				logDir,
				file,
			),
		),

		level,
	)
}

func SetLogger() *zap.Logger {

	cores :=
		[]zapcore.Core{

			createCore(
				zap.LevelEnablerFunc(
					func(
						l zapcore.Level,
					) bool {

						return l ==
							zap.DebugLevel
					},
				),

				debugFileName,
			),

			createCore(
				zap.LevelEnablerFunc(
					func(
						l zapcore.Level,
					) bool {

						return l >=
							zap.InfoLevel &&
							l <
								zap.ErrorLevel
					},
				),

				infoFileName,
			),

			createCore(
				zap.LevelEnablerFunc(
					func(
						l zapcore.Level,
					) bool {

						return l >=
							zap.ErrorLevel
					},
				),

				errorFileName,
			),

			zapcore.NewCore(
				createEncoder(),
				zapcore.Lock(
					os.Stdout,
				),
				zap.InfoLevel,
			),
		}

	return zap.New(
		zapcore.NewTee(
			cores...,
		),

		zap.AddCaller(),

		zap.AddStacktrace(
			zap.ErrorLevel,
		),
	)
}

func AddContext(
	ctx context.Context,
	logger *zap.Logger,
) context.Context {

	return context.WithValue(
		ctx,
		loggerKey,
		logger,
	)
}

func FromContext(
	ctx context.Context,
) *zap.Logger {

	logger, ok :=
		ctx.Value(
			loggerKey,
		).(*zap.Logger)

	if !ok {

		return zap.L()
	}

	return logger
}

func Logger(
	c *gin.Context,
) *zap.Logger {

	return FromContext(
		c.Request.Context(),
	)
}
