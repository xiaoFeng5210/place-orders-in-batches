package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var Logger *zap.Logger

func InitZap() {
	// today := time.Now().Format("2006-01-02")

	// logFileName := fmt.Sprintf("./log/%s.log", today)
	// O_TRUNC 是每次都会清空文件内容
	// fout, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	// if err != nil {
	// 	panic(err)
	// }

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000") //指定时间格式
	encoderConfig.TimeKey = "time"                                                    //默认是ts
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder                           //指定level的显示样式
	l, _ := rotatelogs.New(
		"./log/%Y-%m-%d.log",
		rotatelogs.WithMaxAge(30*24*time.Hour),    // 最长保存30天
		rotatelogs.WithRotationTime(time.Hour*24), // 24小时切割一次
	)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig), //json格式
		zapcore.AddSync(l),                    //指定输出到文件
		zapcore.InfoLevel,                     //设置最低级别
	)
	Logger = zap.New(
		core,
		zap.AddCaller(), // 上报文件名和行号
	)
	Logger = Logger.With(
		zap.Namespace("uber"), //后续的Field都记录在此命名空间中
		//通过zap.String、zap.Int等显式指定类型；fmt.Printf之类的方法大量使用interface{}和反射，性能损失不少
	)
	defer Logger.Sync() //把缓冲里的内容刷入磁盘
}
