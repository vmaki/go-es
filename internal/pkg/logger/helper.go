package logger

import (
	"encoding/json"
	"go-es/global"
	"go.uber.org/zap"
)

// LogIf 当 err != nil 时记录 error 等级的日志
func LogIf(err error) {
	if err != nil {
		global.GLog.Error("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试日志，详尽的程序日志
// 调用示例: logger.Debug("Database", zap.String("sql", sql))
func Debug(moduleName string, fields ...zap.Field) {
	global.GLog.Debug(moduleName, fields...)
}

func Info(moduleName string, fields ...zap.Field) {
	global.GLog.Info(moduleName, fields...)
}

func Warn(moduleName string, fields ...zap.Field) {
	global.GLog.Warn(moduleName, fields...)
}

func Error(moduleName string, fields ...zap.Field) {
	global.GLog.Error(moduleName, fields...)
}

// DebugString 记录一条字符串类型的 debug 日志，调用示例：
// 调用示例: logger.DebugString("SMS", "短信", "123456")
func DebugString(moduleName, name, msg string) {
	global.GLog.Debug(moduleName, zap.String(name, msg))
}

func InfoString(moduleName, name, msg string) {
	global.GLog.Info(moduleName, zap.String(name, msg))
}

func WarnString(moduleName, name, msg string) {
	global.GLog.Warn(moduleName, zap.String(name, msg))
}

func ErrorString(moduleName, name, msg string) {
	global.GLog.Error(moduleName, zap.String(name, msg))
}

// DebugJSON 记录对象类型的 debug 日志，使用 json.Marshal 进行编码。调用示例：
// 调用示例: logger.DebugJSON("Auth", "登录", auth.CurrentUser())
func DebugJSON(moduleName, name string, value interface{}) {
	global.GLog.Debug(moduleName, zap.String(name, jsonString(value)))
}

func InfoJSON(moduleName, name string, value interface{}) {
	global.GLog.Info(moduleName, zap.String(name, jsonString(value)))
}

func WarnJSON(moduleName, name string, value interface{}) {
	global.GLog.Warn(moduleName, zap.String(name, jsonString(value)))
}

func ErrorJSON(moduleName, name string, value interface{}) {
	global.GLog.Error(moduleName, zap.String(name, jsonString(value)))
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		global.GLog.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}

	return string(b)
}
