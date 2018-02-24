package server

import (
	"net/http"
	"time"

	"github.com/xxlixin1993/CacheGo/configure"
	"io"
)

// Initialize http server
func initHttpServer() error {
	host := configure.DefaultString("http.host", "0.0.0.0")
	port := configure.DefaultString("http.port", "12345")
	readTimeout := configure.DefaultInt("http.readTimeout", 3)
	writeTimeout := configure.DefaultInt("http.writeTimeout", 3)
	socketLink := host + ":" + port

	httpServer = &HttpServer{
		host:       host,
		port:       port,
		socketLink: socketLink,
		server: &http.Server{
			Addr:         socketLink,
			Handler:      getServerMux(),
			ReadTimeout:  time.Duration(readTimeout) * time.Second,
			WriteTimeout: time.Duration(writeTimeout) * time.Second,
		},
	}

	return nil
}

// Start http server
func startHttpServer() error {
	initErr := initHttpServer()
	if initErr != nil {
		return initErr
	}

	serveErr := httpServer.server.ListenAndServe()
	if serveErr != nil {
		return serveErr
	}

	return nil
}

// Get a new ServeMux.
func getServerMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", status)
	return mux
}

// Get CacheGo status
func status(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO show status")
}
