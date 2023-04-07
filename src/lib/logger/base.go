package logger

import (
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/**
 * 建立 logger
 * filePath      日誌檔案路徑
 * level         日誌級別
 * maxSize       每個日誌檔案儲存的最大尺寸 單位：M
 * maxBackups    日誌檔案最多儲存多少個備份
 * maxAge        檔案最多儲存多少天
 * compress      是否壓縮
 * serviceName   服務名稱
 * isShowConsole 是否要 show 在 console
 */
func newLogger(filePath string, level zapcore.Level,
	maxSize int, maxBackups int, maxAge int, compress bool,
	serviceName string, isShowConsole bool) *zap.Logger {

	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress, isShowConsole)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.AddCallerSkip(1), zap.Fields(zap.String("serviceName", serviceName)))
}

/**
 * 建立 zapcore
 */
func newCore(filePath string, level zapcore.Level,
	maxSize int, maxBackups int, maxAge int, compress bool, isShowConsole bool) zapcore.Core {
	//日誌檔案路徑配置2
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日誌檔案路徑
		MaxSize:    maxSize,    // 每個日誌檔案儲存的最大尺寸 單位：M
		MaxBackups: maxBackups, // 日誌檔案最多儲存多少個備份
		MaxAge:     maxAge,     // 檔案最多儲存多少天
		Compress:   compress,   // 是否壓縮
	}
	// 設定日誌級別
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用編碼器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小寫編碼器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 時間格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路徑編碼器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 列印到控制檯和檔案
	var writeSyncer zapcore.WriteSyncer
	if isShowConsole {
		writeSyncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		writeSyncer = zapcore.AddSync(&hook)
	}

	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		writeSyncer,
		atomicLevel,
	)
}
