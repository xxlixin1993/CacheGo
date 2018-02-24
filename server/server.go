package server

import (
	"errors"

	"github.com/xxlixin1993/CacheGo/configure"
)

// Start server
func StartServer() error {
	protocolType := configure.DefaultString("server.support", "http")

	switch protocolType {
	case "http":
		return startHttpServer()
	case "tcp":
		return startTcpServer()
	default:
		return errors.New("unknow server.support in configure")
	}

	return nil
}
