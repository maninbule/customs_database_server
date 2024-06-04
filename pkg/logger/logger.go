package logger

import (
	"context"
	"encoding/json"
	"fmt"
	_ "gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"runtime"
	"time"
)

type Level int8

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

type Fields map[string]interface{}

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "debug"
	case LevelInfo:
		return "info"
	case LevelWarn:
		return "warn"
	case LevelError:
		return "error"
	case LevelFatal:
		return "fatal"
	case LevelPanic:
		return "panic"
	default:
		return "error"
	}
}

type Logger struct {
	newLogger *log.Logger
	ctx       context.Context
	fields    Fields
	callers   []string
}

func NewLogger(w io.Writer, prefix string, flag int) *Logger {
	l := log.New(w, prefix, flag)
	return &Logger{newLogger: l}
}

func (l *Logger) clone() *Logger {
	newlog := *l
	return &newlog
}

//WithFields 设置公共字段

func (l *Logger) WithFields(f Fields) *Logger {
	newlog := l.clone()
	if newlog.fields == nil {
		newlog.fields = make(Fields)
	}
	for k, v := range f {
		newlog.fields[k] = v
	}
	return newlog
}

//WithContext 设置日志上下文属性

func (l *Logger) WithContext(ctx context.Context) *Logger {
	newlog := l.clone()
	newlog.ctx = ctx
	return newlog
}

// WithCaller 设置当前某一层(skip) 调用栈的信息(对应函数信息，文件信息，行号)

func (l *Logger) WithCaller(skip int) *Logger {
	newlog := l.clone()
	pc, file, line, ok := runtime.Caller(skip)
	if ok {
		f := runtime.FuncForPC(pc)
		newlog.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return newlog
}

// WithCallersFrames 设置记录整个调用栈信息

func (l *Logger) WithCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	callers := []string{}
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		info := fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function)
		callers = append(callers, info)
		if !more {
			break
		}
	}
	newlog := l.clone()
	newlog.callers = callers
	return newlog
}

// JSONFormat 将要输出的key-value整合到一个map中

func (l *Logger) JSONFormat(level Level, message string) map[string]interface{} {
	data := make(Fields, len(l.fields)+4)
	data["level"] = level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = message
	data["callers"] = l.callers
	if len(l.fields) > 0 {
		for k, v := range l.fields {
			data[k] = v
		}
	}
	return data
}

//Output 将对应level级别内容进行log记录写入磁盘

func (l *Logger) Output(level Level, message string) {
	marshal, _ := json.Marshal(l.JSONFormat(level, message))
	content := string(marshal)
	switch level {
	case LevelDebug:
		l.newLogger.Print(content)
	case LevelInfo:
		l.newLogger.Print(content)
	case LevelWarn:
		l.newLogger.Print(content)
	case LevelError:
		l.newLogger.Print(content)
	case LevelFatal:
		l.newLogger.Fatal(content)
	case LevelPanic:
		l.newLogger.Panic(content)
	}
}

// 各类级别信息级别的日志输出，外部调用

func (l *Logger) Info(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Infof(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Debug(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Warn(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Warnf(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Error(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Fatal(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (l *Logger) Panic(v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprint(v...))
}

func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Output(LevelInfo, fmt.Sprintf(format, v...))
}
