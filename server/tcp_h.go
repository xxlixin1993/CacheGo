package server

import (
	"net"
	"sync"
)

const KTcpServerModuleName = "tcpServerModule"

type TcpServer struct {
	host       string
	port       string
	socketLink string
	listener   net.Listener
	sync.WaitGroup
}