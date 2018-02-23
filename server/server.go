package server

import (
	"fmt"
	"github.com/xxlixin1993/CacheGo/configure"
	"github.com/xxlixin1993/CacheGo/logging"
	"net"
)

// Initialize tcp server
func InitTcpServer() {
	tcpHost := configure.DefaultString("tcp.host", "0.0.0.0")
	tcpPort := configure.DefaultString("tcp.port", "12345")
	socketLink := tcpHost + ":" + tcpPort

	tcpServer = &TcpServer{
		host:       tcpHost,
		port:       tcpPort,
		socketLink: socketLink,
	}
}

// Start tcp server
func StartTcpServer() error {
	var listenErr error
	tcpListener, listenErr = net.Listen("tcp", tcpServer.socketLink)

	if listenErr != nil {
		return listenErr
	}

	defer closeTcpServer()

	for {
		conn, acceptErr := tcpListener.Accept()
		if acceptErr != nil {
			logging.Error("[server] Accept error, msg: ", acceptErr)
		}
		go connHandler(conn)
	}

	return nil
}

func GetTcpServer() *TcpServer {
	return tcpServer
}

func GetTcpListener() net.Listener {
	return tcpListener
}

func connHandler(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Error reading", err.Error())
			//break
			return
		}
		fmt.Printf("Received data: %v", string(buf[:len]))
	}
}

func closeTcpServer() {
	tcpListener.Close()
}
