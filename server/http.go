package server

import (
	"context"
	"io"
	"net/http"
	"time"

	"fmt"
	"github.com/xxlixin1993/CacheGo/configure"
	"github.com/xxlixin1993/CacheGo/distributed"
	"github.com/xxlixin1993/CacheGo/logging"
	"github.com/xxlixin1993/CacheGo/lru"
	"github.com/xxlixin1993/CacheGo/utils"
	"strings"
)

// Initialize http server
func initHttpServer() error {
	host := configure.DefaultString("host", "0.0.0.0")
	port := configure.DefaultString("port", "12345")
	readTimeout := configure.DefaultInt("http.read_timeout", 4)
	writeTimeout := configure.DefaultInt("http.write_timeout", 3)
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

	// graceful exit
	utils.GetExitList().Pop(httpServer)

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
	mux.HandleFunc("/get", get)
	mux.HandleFunc("/set", set)
	mux.HandleFunc("/del", del)
	return mux
}

// Implement ExitInterface
func (h *HttpServer) GetModuleName() string {
	return KHttpServerModuleName
}

// Implement ExitInterface
func (h *HttpServer) Stop() error {
	quitTimeout := configure.DefaultInt("http.quit_timeout", 30)
	ctx, _ := context.WithTimeout(context.Background(), time.Duration(quitTimeout)*time.Second)

	return httpServer.server.Shutdown(ctx)
}

// Get CacheGo status
func status(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "TODO show status")
}

// Get the item in the cache
func get(w http.ResponseWriter, r *http.Request) {
	selfHashNode := configure.DefaultString("host", "127.0.0.1") + ":" +
		configure.DefaultString("port", "12345")

	r.ParseForm()
	key := r.Form.Get("key")

	node := distributed.GetHashRing()
	if node == nil {
		outputHttp(w, "Plz init hash ring first")
	}

	nodeName := node.Get(key)
	if nodeName == "" {
		logging.ErrorF("[http] can not find key(%s)", key)
	}

	if nodeName == selfHashNode {

		// Current node
		if value, ok := lru.LRUCache.Get(key); ok {
			outputHttp(w, value)
		} else {
			outputHttp(w, "")
		}
	} else {
		// Remote node
		cmdUrl := "http://" + nodeName + KGetHttpUrl + "?key=" + key
		response, err := sendHttpRemoteNode(cmdUrl)
		fmt.Println(cmdUrl)
		fmt.Println(response, err)
		if err != nil {
			logging.ErrorF("[http] sendHttpRemoteNode error(%s)", err)
			outputHttp(w, "0")
		}
		outputHttp(w, response)
	}
}

// Set the item in the cache
func set(w http.ResponseWriter, r *http.Request) {
	selfHashNode := configure.DefaultString("host", "127.0.0.1") + ":" +
		configure.DefaultString("port", "12345")

	r.ParseForm()
	key := r.Form.Get("key")
	value := r.Form.Get("value")
	if value = strings.TrimSpace(value); len(value) == 0 {
		outputHttp(w, "Plz input value")
	}

	node := distributed.GetHashRing()
	if node == nil {
		outputHttp(w, "Plz init hash ring first")
	}

	nodeName := node.Get(key)
	if nodeName == "" {
		logging.ErrorF("[http] can not find key(%s)", key)
	}

	if nodeName == selfHashNode {
		// Current node
		if ok := lru.LRUCache.Add(key, value); ok {
			outputHttp(w, "1")
		} else {
			outputHttp(w, "0")
		}
	} else {
		// Remote node
		cmdUrl := "http://" + nodeName + KSetHttpUrl + "?key=" + key + "&value=" + value
		response, err := sendHttpRemoteNode(cmdUrl)
		if err != nil {
			logging.ErrorF("[http] sendHttpRemoteNode error(%s)", err)
			outputHttp(w, "0")
		}
		outputHttp(w, response)
	}
}

// Delete the item in the cache
func del(w http.ResponseWriter, r *http.Request) {
	selfHashNode := configure.DefaultString("host", "127.0.0.1") + ":" +
		configure.DefaultString("port", "12345")
	r.ParseForm()
	key := r.Form.Get("key")
	node := distributed.GetHashRing()

	if node == nil {
		outputHttp(w, "Plz init hash ring first")
	}

	nodeName := node.Get(key)
	if nodeName == "" {
		logging.ErrorF("[http] can not find key(%s)", key)
	}

	if nodeName == selfHashNode {
		// Current node
		if ok := lru.LRUCache.Delete(key); ok {
			outputHttp(w, "1")
		} else {
			outputHttp(w, "0")
		}
	} else {
		// Remote node
		cmdUrl := "http://" + nodeName + KDelHttpUrl + "?key=" + key
		response, err := sendHttpRemoteNode(cmdUrl)
		if err != nil {
			logging.ErrorF("[http] sendHttpRemoteNode error(%s)", err)
			outputHttp(w, "0")
		}
		outputHttp(w, response)
	}
}
