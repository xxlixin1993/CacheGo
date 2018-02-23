package server

import (
	"net"
	"sync"
)

var tcpServer *TcpServer

var tcpListener net.Listener

type TcpServer struct {
	host       string
	port       string
	socketLink string
	sync.WaitGroup
}
