package log

import (
	"os"	
	"time"
	
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log log struct
type Log struct {
    ZapLogger *zap.Logger 
}

var logger *Log

// Initialize function
func Initialize(devMode bool, logFile string) {

	if logger != nil {
		panic("Log already initialized")
	}

	file, _ := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	writerSyncer := zapcore.AddSync(file)		
	encoder := getEncoder()	

	var level zapcore.Level

	if (devMode) {
		level = zapcore.DebugLevel
	} else {
		level = zapcore.WarnLevel
	}
	
	core := zapcore.NewCore(encoder, writerSyncer, level)	
	logger = &Log{ZapLogger: zap.New(core, zap.AddCaller())}	
	
	//logger = log.Sugar() //TODO: ver si tiene sentido utilizar sugar logger
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func (log *Log) getLogger() *zap.Logger {
	return log.ZapLogger
}

// Debug function
func Debug(msg string, fields ...zap.Field) {
	logger.getLogger().Debug(msg, fields...)
}

// Info function
func Info(msg string, fields ...zap.Field) {
	logger.getLogger().Info(msg, fields...)
}

// Warn function
func Warn(msg string, fields ...zap.Field) {
	logger.getLogger().Warn(msg, fields...)
}

// Error function
func Error(msg string, fields ...zap.Field) {
	logger.getLogger().Error(msg, fields...)
}

// Fatal function
func Fatal(msg string, fields ...zap.Field) {
	logger.getLogger().Fatal(msg, fields...)
}

// Panic function
func Panic(msg string, fields ...zap.Field) {
	logger.getLogger().Panic(msg, fields...)
}

// String string field
func String(key string, value string) zap.Field {
	return zap.String(key, value)
}

// Int int field
func Int(key string, value int) zap.Field {
	return zap.Int(key, value)
}

// Duration duration field
func Duration(key string, value time.Duration) zap.Field {
	return zap.Duration(key, value)
}

// Bool bool field
func Bool(key string, value bool) zap.Field {
	return zap.Bool(key, value)
}

// Float float field
func Float(key string, value float64) zap.Field {
	return zap.Float64(key, value)
}

// Time time field
func Time(key string, value time.Time) zap.Field {
	return zap.Time(key, value)
}

// Namespace namespace field
func Namespace(namespace string) zap.Field {
	return zap.Namespace(namespace)
}
	
// Stack stack field
func Stack(key string) zap.Field {
	return zap.Stack(key)
}

// Any any field
func Any(key string, value interface{}) zap.Field {
	return zap.Any(key, value)
}

// Close close function
func Close() {
	logger.getLogger().Info("Logger finalizado")
	logger = nil
}