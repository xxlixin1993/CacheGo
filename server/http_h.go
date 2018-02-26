package server

import (
	"net/http"
)

const KHttpServerModuleName = "httpServerModule"

var httpServer *HttpServer

// HTTP request command url
const (
	KSetHttpUrl = "/set"
	KDelHttpUrl = "/del"
	KGetHttpUrl = "/get"
)

type HttpServer struct {
	host       string
	port       string
	socketLink string
	server     *http.Server
}
