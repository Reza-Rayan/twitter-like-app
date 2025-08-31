package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"time"
)

var Log *zap.Logger

func InitLogger() {
	//	Create logs folder if does not exist
	logDIR := "logs"
	_, err := os.Stat(logDIR)
	if err != nil {
		_ = os.Mkdir(logDIR, os.ModePerm)
	}

	// lumberjack for rotation
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logDIR, "app.log"),
		MaxSize:    100, // 100 MB
		MaxBackups: 10,
		MaxAge:     3, // 3 days
		Compress:   true,
	})

	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "time"
	encoderCfg.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)
	encoderCfg.LevelKey = "level"
	encoderCfg.CallerKey = "caller"
	encoderCfg.MessageKey = "msg"

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg), // Encode JSON format
		w,
		zap.InfoLevel,
	)

	Log = zap.New(core, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
}
