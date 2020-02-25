package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

//////  logger

//var logger *zap.Logger
//
//func main() {
//	logger, _ = zap.NewProduction()
//	defer logger.Sync()
//	Get("http://www.baidu.com")
//	Get("http://www.google.com")
//}
//
//
//func Get(url string) {
//	resp, err := http.Get(url)
//	if err != nil {
//		logger.Error("Error fetching url  ", zap.String("url", url), zap.Error(err))
//		return
//	}
//
//	logger.Info("Success  ", zap.String("url", url))
//	resp.Body.Close()
//}



////// zap logger
//var sugarLogger *zap.SugaredLogger
//
//func main() {
//	logger, _ := zap.NewProduction()
//	sugarLogger = logger.Sugar()
//	defer sugarLogger.Sync()
//	Get("http://www.baidu.com")
//	Get("http://www.google.com")
//}
//
//
//func Get(url string) {
//	sugarLogger.Debugf("Get request for %s", url)
//	resp, err := http.Get(url)
//	if err != nil {
//		sugarLogger.Errorf("Error-----> URL %s : Error = %s", url, err)
//		return
//	}
//	sugarLogger.Infof("Success!!!  for URL %s", url)
//	resp.Body.Close()
//}



/// 自定义
var SugarLogger *zap.SugaredLogger
// 初始化log
func InitLog() {
	level := GetLogLevel()
	writeSyncer, _ := os.Create("./test.log")
	encoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	core := zapcore.NewCore(encoder, writeSyncer, level)
	logger := zap.New(core)
	SugarLogger = logger.Sugar()
	defer SugarLogger.Sync()
}

func GetLogLevel() zapcore.Level {
	level := "debug"		// 从配置文件或用户输入来获取log level
	switch level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	}
	return zapcore.DebugLevel
}