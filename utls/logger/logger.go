package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"regexp"
	"strings"
)

type MyWriter struct {
}

// Write 实现 io.Writer 接口的 Write 方法
func (cw *MyWriter) Write(p []byte) (n int, err error) {
	// 外层函数已加锁

	// 输出前
	compile := regexp.MustCompile(`(^\[\S+])`)
	levelStr := strings.TrimFunc(
		compile.FindString(string(p)), func(r rune) bool {
			return r == '[' || r == ']'
		})
	outStr := string(p)[len(levelStr)+5:]
	beforeOut(levelStr, &outStr)

	// 调用底层输出
	n, err = os.Stdout.Write([]byte(outStr))

	// 输出后
	afterOut()

	return n, err
}

type MyFormatter struct{}

// Format 格式化日志条目
func (f *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {

	// 构建日志格式
	level := strings.ToUpper(entry.Level.String())

	message := entry.Message

	// 构建输出字符串
	formatted := "[" + level + "] " + ": " + message + "\n"
	return []byte(formatted), nil
}

var myLogger = logrus.New()

// 仅显示时间，但写入文件
func init() {
	myLogger.SetFormatter(&MyFormatter{})
	myLogger.SetOutput(&MyWriter{})
	myLogger.SetLevel(logrus.TraceLevel)
}

func Debug(v ...any) {
	myLogger.Debug(v...)
}

func Info(v ...any) {
	myLogger.Info(v...)
}

func Warning(v ...any) {
	myLogger.Warning(v...)
}

func Error(v ...any) {
	myLogger.Error(v...)
	//os.Exit(1)
}

func Trace(v ...any) {
	myLogger.Trace(v...)
}

// DebugF 带格式化的调试日志
func DebugF(format string, v ...any) {
	myLogger.Debugf(format, v...)
}

// InfoF 带格式化的信息日志
func InfoF(format string, v ...any) {
	myLogger.Infof(format, v...)

}

// WarningF 带格式化的警告日志
func WarningF(format string, v ...any) {
	myLogger.Warningf(format, v...)
}

// ErrorF 带格式化的错误日志
func ErrorF(format string, v ...any) {
	myLogger.Errorf(format, v...)
	//os.Exit(1)
}

// TraceF 带格式化的追踪日志
func TraceF(format string, v ...any) {
	myLogger.Tracef(format, v...)
}
