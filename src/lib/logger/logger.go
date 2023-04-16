package logger

import (
	"fmt"
	"openaigo/config"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type ILogger interface {
	Field(key string, val interface{}) zap.Field
	ApplicationDebug(msg string, fields ...zapcore.Field)
	ApplicationInfo(msg string, fields ...zapcore.Field)
	ApplicationWarn(msg string, fields ...zapcore.Field)
	ApplicationError(msg string, fields ...zapcore.Field)
	ApplicationErrorSimple(msg string, err error)
	BusinessDebug(msg string, fields ...zapcore.Field)
	BusinessInfo(msg string, fields ...zapcore.Field)
}

type ZapLogger struct {
	applicationMutex sync.Mutex
	Application      *zap.Logger
	businessMutex    sync.Mutex
	Business         *zap.Logger
}

func New(loggerConfig config.Logger) *ZapLogger {
	hostName, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	filePath := loggerConfig.FilePath
	level := getZapLevel(loggerConfig.Level)
	maxSize := loggerConfig.MaxSize
	maxBackups := loggerConfig.MaxBackups
	maxAge := loggerConfig.MaxAge
	compress := loggerConfig.Compress
	serviceName := loggerConfig.ServiceName
	isShowConsole := loggerConfig.IsShowConsole

	applicationFilePath := fmt.Sprintf("%vapplication/%v/%v_%v_log.log", filePath, serviceName, serviceName, hostName)
	businessFilePath := fmt.Sprintf("%vbusiness/%v/%v_%v_log.log", filePath, serviceName, serviceName, hostName)
	createLogFileIfNotExist(applicationFilePath)
	createLogFileIfNotExist(businessFilePath)
	return &ZapLogger{
		Application: newLogger(
			applicationFilePath,
			level,
			maxSize,
			maxBackups,
			maxAge,
			compress,
			serviceName,
			isShowConsole),
		Business: newLogger(
			businessFilePath,
			level,
			maxSize,
			maxBackups,
			maxAge,
			compress,
			serviceName,
			isShowConsole),
	}
}

// getZapLevel Mapping zap level
func getZapLevel(level string) zapcore.Level {
	switch l := strings.ToLower(level); l {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func createLogFileIfNotExist(file string) {
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, os.ModePerm)
	}
}

// Field make zepField
func (l *ZapLogger) Field(key string, val interface{}) zap.Field {
	return zap.Any(key, val)
}

// ApplicationDebug add application debug log
func (l *ZapLogger) ApplicationDebug(msg string, fields ...zapcore.Field) {
	defer l.applicationMutex.Unlock()
	l.applicationMutex.Lock()
	l.Application.Debug(msg, fields...)
}

// ApplicationInfo add application info log
func (l *ZapLogger) ApplicationInfo(msg string, fields ...zapcore.Field) {
	defer l.applicationMutex.Unlock()
	l.applicationMutex.Lock()
	l.Application.Info(msg, fields...)
}

// ApplicationWarn add application warn log
func (l *ZapLogger) ApplicationWarn(msg string, fields ...zapcore.Field) {
	defer l.applicationMutex.Unlock()
	l.applicationMutex.Lock()
	l.Application.Warn(msg, fields...)
}

// ApplicationError add application error log
func (l *ZapLogger) ApplicationError(msg string, fields ...zapcore.Field) {
	defer l.applicationMutex.Unlock()
	l.applicationMutex.Lock()
	l.Application.Error(msg, fields...)
}

// ApplicationErrorSimple add application error log
func (l *ZapLogger) ApplicationErrorSimple(msg string, err error) {
	defer l.applicationMutex.Unlock()
	l.applicationMutex.Lock()
	l.Application.Error(msg, zap.Error(err))
}

// BusinessDebug add business debug log
func (l *ZapLogger) BusinessDebug(msg string, fields ...zapcore.Field) {
	defer l.businessMutex.Unlock()
	l.businessMutex.Lock()
	l.Business.Debug(msg, fields...)
}

// BusinessInfo add business info log
func (l *ZapLogger) BusinessInfo(msg string, fields ...zapcore.Field) {
	defer l.businessMutex.Unlock()
	l.businessMutex.Lock()
	l.Business.Info(msg, fields...)
}
