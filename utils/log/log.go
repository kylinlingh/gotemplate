package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// 使用教程：https://www.liwenzhou.com/posts/Go/zap/

var GlobalLogger *zap.SugaredLogger


func init() {
	var coreArray []zapcore.Core
	encoder := getEncoder()
	consoleCore := zapcore.NewCore(encoder, getConsoleWriter(),zapcore.DebugLevel)

	// 只输出到控制台
	coreArray = append(coreArray, consoleCore)

	// 同时输出到控制台与文件，如果文件不存在，会自动创建，不需要特别处理
	//fileCore := zapcore.NewCore(encoder, getFileWriter("./test.log"), zapcore.DebugLevel)
	//coreArray = append(coreArray, consoleCore, fileCore)

	logger := zap.New(zapcore.NewTee(coreArray...)).WithOptions(zap.AddCaller()) // 记录函数所在的位置
	GlobalLogger = logger.Sugar()
}

func getEncoder() zapcore.Encoder {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) { // 输出可阅读形式的时间戳
		enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
	}
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 将日志级别以缩写形式输出，并使用不同颜色

	// 以json格式输出
	//return zapcore.NewJSONEncoder(encoderConfig)

	// 以普通格式输出
	return zapcore.NewConsoleEncoder(encoderConfig)
}


func getFileWriter(filePath string) zapcore.WriteSyncer {
	// 设置日志回滚
	hook := lumberjack.Logger{
		Filename:   filePath, // 自动追加文件末尾
		MaxSize:    100, // megabytes
		MaxBackups: 30,
		MaxAge:     30,   //days
		Compress:   true, // disabled by default
	}

	return zapcore.AddSync(&hook)
}

func getConsoleWriter() zapcore.WriteSyncer{
	console := zapcore.Lock(os.Stdout)
	return console
}



