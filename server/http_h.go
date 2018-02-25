package server

import (
	"net/http"
)

const KHttpServerModuleName = "httpServerModule"

var httpServer *HttpServer

type HttpServer struct {
	host       string
	port       string
	socketLink string
	server     *http.Server
}
