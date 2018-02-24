package configure

import (
	"sync"
)

var appConfig *Config

var (
	DefaultSelection  = "local"
	DefaultComment    = []byte{'#'}
	DefaultCommentSem = []byte{';'}
)

type Config struct {
	// map is not safe.
	sync.RWMutex

	// Section:key=value
	data map[string]map[string]string
}

// Error code
const (
	KInitConfigError = iota + 1
	KInitLogError
	KInitTcpServerError
	KInitSeverError
)

// Error message
const (
	KUnknownTypeMsg = "unknown type"
)
