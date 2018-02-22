package logging

import (
	"errors"
	"fmt"
	"path"
	"runtime"

	"github.com/xxlixin1993/CacheGo/configure"
	"github.com/xxlixin1993/CacheGo/utils"
)

// Initialize Log
func InitLog() error {
	outputType := configure.DefaultString("log.type", KOutputStdout)
	level := configure.DefaultInt("log.level", KLevelDebug)

	logger, err := getLogger(outputType, level)
	if err != nil {
		return err
	}

	logger.Add(1)
	go logger.Run()

	SetLogger(logger)

	return err
}

func getLogger(outputType string, level int) (*LogBase, error) {
	switch outputType {
	case KOutputStdout:
		return &LogBase{
			handle:  NewStdoutLog(),
			message: make(chan []byte, 1000),
			skip:    3,
			level:   level,
		}, nil
	case KOutputFile:
		// TODO
		return nil, errors.New("TODO not supported")
	default:
		return nil, errors.New(configure.KUnknownTypeMsg)
	}
}

func WaitLog() {
	close(loggerInstance.message)
	loggerInstance.Wait()
}

func SetLogger(logger *LogBase) {
	loggerInstance = logger
}

func GetLogger() *LogBase {
	return loggerInstance
}

func Debug(args ...interface{}) {
	msg := fmt.Sprint(args...)
	GetLogger().Output(KLevelDebug, msg)
}

func Trace(args ...interface{}) {
	msg := fmt.Sprint(args...)
	GetLogger().Output(KLevelTrace, msg)
}

func Info(args ...interface{}) {
	msg := fmt.Sprint(args...)
	GetLogger().Output(KLevelInfo, msg)
}

func Notice(args ...interface{}) {
	msg := fmt.Sprint(args...)
	GetLogger().Output(KLevelNotice, msg)
}

func Warnning(args ...interface{}) {
	msg := fmt.Sprint(args...)
	GetLogger().Output(KLevelWarnning, msg)
}

func Error(args ...interface{}) {
	msg := fmt.Sprint(args...)
	GetLogger().Output(KLevelError, msg)
}

func Fatal(args ...interface{}) {
	msg := fmt.Sprint(args...)
	GetLogger().Output(KLevelFatal, msg)
}

func (l *LogBase) Run() {
	for {
		msg, ok := <-l.message
		if !ok {
			l.Done()
			break
		}
		err := l.handle.OutputLogMsg(msg)
		if err != nil {
			fmt.Printf("Log: Output handle fail, err:%v\n", err.Error())
		}
	}
}

func (l *LogBase) Output(nowLevel int, msg string) {
	now := utils.GetMicTimeFormat()

	l.mu.Lock()
	defer l.mu.Unlock()

	if nowLevel <= l.level {
		_, file, line, ok := runtime.Caller(l.skip)
		if !ok {
			file = "???"
			line = 0
		}
		_, filename := path.Split(file)
		msg = fmt.Sprintf("[%s] [%s %s:%d] %s\n", LevelName[nowLevel], now, filename, line, msg)
	}

	l.message <- []byte(msg)
}
