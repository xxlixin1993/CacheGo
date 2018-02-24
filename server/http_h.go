package server

import (
	"net"
	"net/http"
	"sync"
)

var httpServer *HttpServer

type TcpServer struct {
	host       string
	port       string
	socketLink string
	listener   net.Listener
	sync.WaitGroup
}

type HttpServer struct {
	host       string
	port       string
	socketLink string
	server     *http.Server
}
