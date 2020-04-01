package mlog

import (
	"github.com/kataras/iris/context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"orange_message_service/app/common"
	"orange_message_service/app/components/config"
	"os"
	"time"
)

var logger *zap.Logger

func Init() {
	config := config.GetConfig()
	cfg := zap.NewProductionConfig()

	hook := lumberjack.Logger{
		Filename:   config.GetString("log.filename"),
		MaxSize:    config.GetInt("log.max_size"),
		MaxBackups: config.GetInt("log.max_backups"),
		MaxAge:     config.GetInt("log.max_age"),
		Compress:   false,
	}

	isDev := config.GetString("mode") == common.ENV_DEV
	var ws zapcore.WriteSyncer
	if isDev {
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook))
	}
	cfg.Encoding = "json"
	cfg.EncoderConfig.EncodeTime = syslogTimeEncoder
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	atomicLevel := zap.NewAtomicLevel()
	if isDev {
		atomicLevel.SetLevel(zap.DebugLevel)
	} else {
		atomicLevel.SetLevel(zap.InfoLevel)
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), ws, atomicLevel)

	logger = zap.New(core)
}

func GetLogger() *zap.Logger {
	return logger
}

func syslogTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func WarnCommonError(ctx context.Context, err *common.ComError) {
	Warnf(ctx, "common error.||errno=%d||msg=%d", err.Code, err.Msg)
}

func ZapInfo(args ...interface{}) {
	logger.Sugar().Info(args...)
}

func ZapInfof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func Info(ctx context.Context, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Info(args...)
}

func Infof(ctx context.Context, template string, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Infof(template, args...)
}

func Warn(ctx context.Context, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Warn(args...)
}

func Warnf(ctx context.Context, template string, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Warnf(template, args...)
}

func WarnfWithFields(commonFields []interface{}, template string, args ...interface{}) {
	logger.Sugar().With(commonFields...).Warnf(template, args...)
}

func Error(ctx context.Context, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Error(args...)
}

func Errorf(ctx context.Context, template string, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Errorf(template, args...)
}

// Debug 是由程序自己控制的，通过上游的调用链来开启
func Debug(ctx context.Context, args ...interface{}) {
	if ctx.GetHeader(common.HTTP_HEADER_MJ_TRACE_ID) != "" {
		logger.Sugar().With(GetComFields(ctx)...).Info(args...)
	} else {
		logger.Sugar().With(GetComFields(ctx)...).Debug(args...)
	}
}

// Debug 是由程序自己控制的，通过上游的调用链来开启
func Debugf(ctx context.Context, template string, args ...interface{}) {
	if ctx.GetHeader(common.HTTP_HEADER_MJ_TRACE_ID) != "" {
		logger.Sugar().With(GetComFields(ctx)...).Infof(template, args...)
	} else {
		logger.Sugar().With(GetComFields(ctx)...).Debugf(template, args...)
	}
}

func Fatal(ctx context.Context, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Fatal(args...)
}

func Fatalf(ctx context.Context, template string, args ...interface{}) {
	logger.Sugar().With(GetComFields(ctx)...).Fatalf(template, args...)
}

func GetComFields(ctx context.Context) []interface{} {
	f, _ := ctx.Values().Get(common.COMMON_LOG_FIELD_KEY).(common.CommonLogFields)
	return []interface{}{
		"ip", f.IP,
		"method", f.Method,
		"path", f.Path,
		"trace_id", f.TraceID,
		"header", f.Header,
	}
}
