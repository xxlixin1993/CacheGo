package server

import (
	"net"
	"net/http"
	"sync"
)

const KHttpServerModuleName = "httpServerModule"

var httpServer *HttpServer

type HttpServer struct {
	host       string
	port       string
	socketLink string
	server     *http.Server
}
