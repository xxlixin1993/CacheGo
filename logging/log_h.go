package logging

import "sync"

// Log message level
const (
	KLevelFatal    = iota
	KLevelError
	KLevelWarnning
	KLevelNotice
	KLevelInfo
	KLevelTrace
	KLevelDebug
)

// Log output message level abbreviation
var LevelName = [7]string{"F", "E", "W", "N", "I", "T", "D"}

// Log instance
var loggerInstance *LogBase

// Log output type
const (
	KOutputFile   = "file"
	KOutputStdout = "stdout"
)

// Log interface. Need to be implemented when you want to extend.
type ILog interface {
	// Initialize Logger
	Init(config interface{}) error

	// Output message to log
	OutputLogMsg(msg []byte) error
}

// Log core program
type LogBase struct {
	mu      sync.Mutex
	sync.WaitGroup
	handle  ILog
	message chan []byte
	skip    int
	level   int
}
