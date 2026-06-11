package logger

import (
	"app/config"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 默认 (未匹配时使用)
var DefaultLoggers *zap.Logger

type LoggerManager struct {
	Access *zap.Logger
	App    *zap.Logger
	Error  *zap.Logger
}

// NewLoggerManager 创建日志管理器 (access\app\error)
func NewLoggerManager(cfg config.LoggerConfig) *LoggerManager {
	lm := &LoggerManager{}

	levelMap := map[string]zapcore.Level{
		"app":    zapcore.InfoLevel,
		"error":  zapcore.ErrorLevel,
		"access": zapcore.InfoLevel,
	}

	for _, biz := range cfg.Business {
		level := zapcore.InfoLevel
		if lv, ok := levelMap[biz.Name]; ok {
			level = lv
		}
		logger := NewLogger(biz.Path, level, cfg)

		switch biz.Name {
		case "app":
			lm.App = logger
		case "error":
			lm.Error = logger
		case "access":
			lm.Access = logger
		}
	}
	return lm
}

func NewLogger(path string, level zapcore.Level, cfg config.LoggerConfig) *zap.Logger {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path,           // 日志文件路径
		MaxSize:    cfg.MaxSize,    // 每个日志文件最大100MB，超过则切割
		MaxBackups: cfg.MaxBackups, // 保留最近7个备份文件
		MaxAge:     cfg.MaxAge,     // 日志文件最长保留30天
		Compress:   cfg.Compress,   // 是否压缩旧日志文件
		LocalTime:  true,           // 是否使用本地时间
	}

	//  异步写入器（降低日志 I/O 对业务的影响）
	writeSyncer := zapcore.AddSync(lumberJackLogger)
	bufferedWS := &zapcore.BufferedWriteSyncer{
		WS:            writeSyncer,
		Size:          256 * 1024, // 256KB 缓冲
		FlushInterval: 30 * time.Second,
	}

	//  动态级别控制（支持运行时通过 SetLevel 修改）
	atomicLevel := zap.NewAtomicLevelAt(level) // 生产环境默认为 Info
	atomicLevel.SetLevel(level)

	// 创建 Core（可选：同时输出到控制台）
	var core zapcore.Core
	// 只输出到文件（生产环境推荐）
	core = zapcore.NewCore(
		zapcore.NewJSONEncoder(getEncoderConfig()),
		bufferedWS,
		atomicLevel,
	)

	// 构建 Logger，添加扩展选项
	Log := zap.New(core,
		zap.AddCaller(),                       // 显示调用者信息
		zap.AddCallerSkip(1),                  // 跳过封装层（如果有）
		zap.AddStacktrace(zapcore.ErrorLevel), // Error 级别以上输出堆栈
	)
	return Log
}

func getEncoderConfig() zapcore.EncoderConfig {
	//  编码器配置（生产环境使用 JSON 格式）
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "@timestamp", // 时间戳键名
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder, // info / error
		EncodeTime:     zapcore.ISO8601TimeEncoder,    // 2024-01-15T10:30:00.000Z
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder, // package/file.go:42
	}
	return encoderConfig
}
